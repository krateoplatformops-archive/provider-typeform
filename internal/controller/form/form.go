package form

import (
	"context"
	"fmt"
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/krateoplatformops/provider-typeform/apis/form/v1alpha1"
	tfc "github.com/krateoplatformops/provider-typeform/internal/clients/config"
	"github.com/krateoplatformops/provider-typeform/internal/clients/typeform"
	"github.com/krateoplatformops/provider-typeform/internal/controller/form/support"
	"github.com/krateoplatformops/provider-typeform/internal/notifications"
)

const (
	errInvalidCR     = "managed resource is not a Form"
	errClientOpts    = "cannot get client options"
	errNewClient     = "cannot create typeform API client"
	errCheckUpToDate = "cannot determine if resource is up to date"

	reasonCannotCreate = "CannotCreateForm"
	reasonCreated      = "CreatedForm"
	reasonUpdated      = "UpdatedForm"
	reasonCannotUpdate = "CannotUpdateForm"
	reasonDeleted      = "DeletedForm"
	reasonCannotDelete = "CannotDeleteForm"

	serviceName = "provider-typeform"

	annotationKeyFormID = "krateo.io/typeform-id"
)

// Setup adds a controller that reconciles Server managed resources.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.FormGroupKind)

	cps := []managed.ConnectionPublisher{
		managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme()),
	}

	log := o.Logger.WithValues("controller", name)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.FormGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:       mgr.GetClient(),
			log:        log,
			notifierFn: notifications.NewNotifier,
			clientFn:   typeform.NewClient,
		}),
		managed.WithLogger(log),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.Form{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// A connector is expected to produce an ExternalClient
// when its Connect method is called.
type connector struct {
	kube       client.Client
	log        logging.Logger
	notifierFn func(string) *notifications.Notifier
	clientFn   func(opts typeform.ClientOpts) typeform.Client
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	opts, err := tfc.GetConfig(ctx, c.kube, mg)
	if err != nil {
		return nil, errors.Wrap(err, errClientOpts)
	}

	return &external{
		client:   c.clientFn(opts.ClientOpts),
		log:      c.log,
		notifier: c.notifierFn(opts.LogServiceUrl),
	}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	client   typeform.Client
	log      logging.Logger
	notifier *notifications.Notifier
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Form)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errInvalidCR)
	}

	formID := meta.GetExternalName(cr)
	if formID == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	e.log.Debug("Looking for Form", "id", formID)

	instance, err := e.client.GetForm(ctx, formID)
	if err != nil {
		var reqErrs *typeform.Errors
		if !errors.As(err, &reqErrs) {
			return managed.ExternalObservation{}, err
		}

		if !strings.HasSuffix(reqErrs.Code, "NOT_FOUND") {
			return managed.ExternalObservation{}, err
		}
	}
	if instance == nil {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	cr.Status.AtProvider = support.GenerateObservation(instance)
	cr.Status.SetConditions(xpv1.Available())

	e.log.Info("Form already exists",
		"id", instance.ID,
		"displayUrl", cr.Status.AtProvider.DisplayURL)

	upToDate, err := support.IsUpToDate(e.log, &cr.Spec.ForProvider, instance)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errCheckUpToDate)
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: upToDate,
	}, err
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Form)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errInvalidCR)
	}

	deploymentId := mg.GetLabels()["deploymentId"]

	cr.Status.SetConditions(xpv1.Creating())

	spec := cr.Spec.ForProvider.DeepCopy()

	val := support.FromFormParams(spec)
	res, err := e.client.CreateForm(ctx, val)
	if err != nil {
		support.Notify(e.log, e.notifier, notifications.Error(
			notifications.Opts{
				ServiceName:  serviceName,
				DeploymentId: deploymentId,
				Reason:       reasonCannotCreate,
				Message:      err.Error(),
			}))

		return managed.ExternalCreation{}, err
	}

	e.log.Info("Form successfully created", "id", res.ID, "deploymentId", deploymentId)

	support.Notify(e.log, e.notifier, notifications.Info(
		notifications.Opts{
			ServiceName:  serviceName,
			DeploymentId: deploymentId,
			Reason:       reasonCreated,
			Message:      fmt.Sprintf("Form created successfully (id: %s, displayUrl: %s)", res.ID, res.Links.Display),
		}))

	meta.SetExternalName(cr, res.ID)

	return managed.ExternalCreation{ExternalNameAssigned: true}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	// NOOP
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.Form)
	if !ok {
		return errors.New(errInvalidCR)
	}

	deploymentId := mg.GetLabels()["deploymentId"]

	cr.Status.SetConditions(xpv1.Deleting())

	formID := cr.Status.AtProvider.ID

	e.log.Info("Deleting Form", "id", formID, "deploymentId", deploymentId)

	err := e.client.DeleteForm(ctx, formID)
	if err != nil {
		support.Notify(e.log, e.notifier, notifications.Error(
			notifications.Opts{
				ServiceName:  serviceName,
				DeploymentId: deploymentId,
				Reason:       reasonCannotDelete,
				Message:      err.Error(),
			}))

		return err
	}

	support.Notify(e.log, e.notifier, notifications.Info(
		notifications.Opts{
			ServiceName:  serviceName,
			DeploymentId: deploymentId,
			Reason:       reasonDeleted,
			Message:      fmt.Sprintf("Form successfully deleted (id: %s) ", formID),
		}))

	return err
}

// getFormID returns the form identifier annotation value on the resource.
func getFormID(o metav1.Object) string {
	return o.GetAnnotations()[annotationKeyFormID]
}

// setFormID sets the form identifier annotation of the resource.
func setFormID(o metav1.Object, id string) {
	meta.AddAnnotations(o, map[string]string{annotationKeyFormID: id})
}
