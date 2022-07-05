package config

import (
	"context"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/krateoplatformops/provider-typeform/apis/v1alpha1"
	"github.com/krateoplatformops/provider-typeform/internal/clients"
	"github.com/krateoplatformops/provider-typeform/internal/clients/typeform"
	"github.com/krateoplatformops/provider-typeform/internal/helpers"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	errProviderConfigRefNotGiven = "providerConfig reference is not given"
	errProviderConfigRefNotFound = "cannot get referenced provider config"
	errPasswordSecret            = "cannot get API token from secret"
	errProviderConfigTrack       = "cannot track providerConfig usage"
)

type Config struct {
	typeform.ClientOpts
	LogServiceUrl string
}

// GetConfig to produce a compute instance service clients options.
func GetConfig(ctx context.Context, crc client.Client, mg resource.Managed) (*Config, error) {
	if mg.GetProviderConfigReference() == nil {
		return nil, errors.New(errProviderConfigRefNotGiven)
	}

	pc := &v1alpha1.ProviderConfig{}
	err := crc.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc)
	if err != nil {
		return nil, errors.Wrap(err, errProviderConfigRefNotFound)
	}

	err = resource.NewProviderConfigUsageTracker(crc, &v1alpha1.ProviderConfigUsage{}).Track(ctx, mg)
	if err != nil {
		return nil, errors.Wrap(err, errProviderConfigTrack)
	}

	token, err := getTypeFormApiToken(ctx, crc, pc)
	if err != nil {
		return nil, errors.Wrap(err, errPasswordSecret)
	}

	opts := typeform.ClientOpts{
		Token: token,
	}

	if helpers.BoolValue(pc.Spec.Verbose) {
		opts.HTTPClient = clients.TracerHTTPClient()
	}

	return &Config{opts, helpers.StringValue(pc.Spec.LogServiceUrl)}, nil
}

// getTypeFormApiToken returns the typeform API personal access token stored in a secret.
func getTypeFormApiToken(ctx context.Context, crc client.Client, pc *v1alpha1.ProviderConfig) (string, error) {
	if s := pc.Spec.Credentials.Source; s != xpv1.CredentialsSourceSecret {
		return "", fmt.Errorf("credentials source %s is not currently supported", s)
	}

	csr := pc.Spec.Credentials.SecretRef
	if csr == nil {
		return "", fmt.Errorf("no credentials secret referenced")
	}

	return helpers.GetSecret(ctx, crc, csr.DeepCopy())
}
