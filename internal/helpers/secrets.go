package helpers

import (
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

func GetSecret(ctx context.Context, kube client.Client, ref *xpv1.SecretKeySelector) (string, error) {
	if ref == nil {
		return "", errors.New("no credentials secret referenced")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	sec := &corev1.Secret{}

	if err := kube.Get(ctx, key, sec); err != nil {
		return "", errors.Wrapf(err, "cannot get %s secret", ref.Name)
	}

	return string(sec.Data[ref.Key]), nil
}
