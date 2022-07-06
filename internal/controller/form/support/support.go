package support

import (
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/krateoplatformops/provider-typeform/apis/form/v1alpha1"
	"github.com/krateoplatformops/provider-typeform/internal/clients/typeform"
	"github.com/krateoplatformops/provider-typeform/internal/helpers"
	"github.com/krateoplatformops/provider-typeform/internal/notifications"
)

// IsUpToDate checks whether current state is up-to-date compared to the given set of parameters.
func IsUpToDate(log logging.Logger, in *v1alpha1.FormParams, observed *typeform.Form) (bool, error) {
	desired := FromFormParams(in.DeepCopy())

	diff := (cmp.Diff(desired, observed,
		cmpopts.EquateEmpty(),
		cmp.Comparer(func(a, b *bool) bool {
			if a == nil {
				return (b == nil) || (*b == false)
			}

			if b == nil {
				return (a == nil) || (*a == false)
			}

			return *a == *b
		}),
		cmpopts.SortSlices(func(a, b typeform.Field) bool {
			return a.Ref < b.Ref
		}),
		cmpopts.SortSlices(func(a, b string) bool {
			return a < b
		}),
		cmpopts.IgnoreFields(typeform.Form{}, "ID", "Links"),
		cmpopts.IgnoreFields(typeform.Field{}, "ID"),
	))

	if diff != "" {
		fmt.Printf("\n\n%s\n\n", diff)
		log.Info("IsUpToDate", "Diff", diff)
		return false, nil
	}

	return true, nil
}

// ToFormParameters generates the resource (kind: Form) parameters from a typeform form object.
func ToFormParameters(in *typeform.Form) *v1alpha1.FormParams {
	res := &v1alpha1.FormParams{
		Title:  in.Title,
		Fields: make([]v1alpha1.Field, len(in.Fields)),
	}

	for i, el := range in.Fields {
		res.Fields[i].Type = el.Type
		res.Fields[i].Title = el.Title

		res.Fields[i].Properties.AllowMultipleSelection = el.Properties.AllowMultipleSelection
		res.Fields[i].Properties.AllowOtherChoice = el.Properties.AllowOtherChoice
		res.Fields[i].Properties.AlphabeticalOrder = el.Properties.AlphabeticalOrder
		res.Fields[i].Properties.Description = el.Properties.Description
		res.Fields[i].Properties.Steps = el.Properties.Steps
		res.Fields[i].Properties.Shape = el.Properties.Shape
		res.Fields[i].Properties.Choices = make([]string, len(el.Properties.Choices))
		for j, it := range el.Properties.Choices {
			res.Fields[i].Properties.Choices[j] = it.Label
		}

		if el.Validations != nil {
			res.Fields[i].Validations = &v1alpha1.Validations{}
			res.Fields[i].Validations.MaxLength = el.Validations.MaxLength
			res.Fields[i].Validations.MaxSelection = el.Validations.MaxSelection
			res.Fields[i].Validations.MinSelection = el.Validations.MinSelection
			res.Fields[i].Validations.Required = helpers.LateInitializeBool(el.Validations.Required, false)
		}

		if el.Layout != nil {
			res.Fields[i].Layout = &v1alpha1.Layout{}
			res.Fields[i].Layout.Type = el.Layout.Type
			res.Fields[i].Layout.Placement = el.Layout.Placement
			res.Fields[i].Layout.Attachment = v1alpha1.Attachment{}
			res.Fields[i].Layout.Attachment.Href = el.Layout.Attachment.Href
			res.Fields[i].Layout.Attachment.Type = el.Layout.Attachment.Type
			res.Fields[i].Layout.Attachment.Scale = el.Layout.Attachment.Scale
		}
	}

	return res
}

func FromFormParams(in *v1alpha1.FormParams) *typeform.Form {
	res := &typeform.Form{
		Title:  in.Title,
		Type:   "form",
		Fields: make([]typeform.Field, len(in.Fields)),
	}

	counters := map[string]int{}

	for i, el := range in.Fields {
		n := counters[el.Type]
		counters[el.Type] = n + 1

		res.Fields[i].Type = el.Type
		res.Fields[i].Title = el.Title
		res.Fields[i].Ref = fmt.Sprintf("%s-%d", strings.ReplaceAll(el.Type, "_", "-"), counters[el.Type])

		res.Fields[i].Properties.AllowMultipleSelection = el.Properties.AllowMultipleSelection
		res.Fields[i].Properties.AllowOtherChoice = el.Properties.AllowOtherChoice
		res.Fields[i].Properties.AlphabeticalOrder = el.Properties.AlphabeticalOrder
		res.Fields[i].Properties.Description = el.Properties.Description
		res.Fields[i].Properties.Steps = el.Properties.Steps
		res.Fields[i].Properties.Shape = el.Properties.Shape
		res.Fields[i].Properties.Choices = make([]typeform.Choice, len(el.Properties.Choices))
		for j, it := range el.Properties.Choices {
			res.Fields[i].Properties.Choices[j] = typeform.Choice{Label: it}
		}

		if el.Validations != nil {
			res.Fields[i].Validations = &typeform.Validations{}
			res.Fields[i].Validations.MaxLength = el.Validations.MaxLength
			res.Fields[i].Validations.MaxSelection = el.Validations.MaxSelection
			res.Fields[i].Validations.MinSelection = el.Validations.MinSelection
			res.Fields[i].Validations.Required = el.Validations.Required
		}

		if el.Layout != nil {
			res.Fields[i].Layout = &typeform.Layout{}
			res.Fields[i].Layout.Type = el.Layout.Type
			res.Fields[i].Layout.Placement = el.Layout.Placement
			res.Fields[i].Layout.Attachment = typeform.Attachment{}
			res.Fields[i].Layout.Attachment.Href = el.Layout.Attachment.Href
			res.Fields[i].Layout.Attachment.Type = el.Layout.Attachment.Type
			res.Fields[i].Layout.Attachment.Scale = el.Layout.Attachment.Scale
		}
	}

	return res
}

// GenerateObservation produces FormObservation object from a typeform.Form object.
func GenerateObservation(src *typeform.Form) v1alpha1.FormObservation {
	return v1alpha1.FormObservation{
		ID:         src.ID,
		DisplayURL: src.Links.Display,
	}
}

func Notify(log logging.Logger, notifier *notifications.Notifier, msg notifications.Notification) {
	go func() {
		err := notifier.Send(msg)
		if err != nil {
			log.Info("Unable to send notification", "deploymentId", msg.TransactionId, "error", err.Error())
		}
	}()
}
