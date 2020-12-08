package client

import (
	"github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1"
	croType "github.com/integr8ly/cloud-resource-operator/pkg/apis/integreatly/v1alpha1/types"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// define a function that allows us to perform modification logic on the
// custom resource (e.g. setting owner refs) before creating or updating it
type modifyResourceFunc func(cr metav1.Object) error

// ReconcileBlobStorage creates or updates a blob storage custom resource
func (c CloudResourceService) ReconcileBlobStorage(productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.BlobStorage, error) {
	bs := &v1alpha1.BlobStorage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels: map[string]string{
				"productName": productName,
			},
		},
	}

	// execute logic to modify the resource before creation
	// e.g. add owner refs
	if modifyFunc != nil {
		err := modifyFunc(bs)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute modification function on resource %s", name)
		}
	}

	// Create or update the resource
	_, err := controllerutil.CreateOrUpdate(c.Ctx, c.Client, bs, func() error {
		bs.Spec.Type = deploymentType
		bs.Spec.Tier = tier
		bs.Spec.SecretRef = &croType.SecretRef{
			Name:      secretName,
			Namespace: secretNs,
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to reconcile blob storage request for %s", name)
	}

	return bs, nil
}

// ReconcilePostgres creates or updates a postgres custom resource
func (c CloudResourceService) ReconcilePostgres(productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.Postgres, error) {
	pg := &v1alpha1.Postgres{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels: map[string]string{
				"productName": productName,
			},
		},
	}

	// execute logic to modify the resource before creation
	// e.g. add owner refs
	if modifyFunc != nil {
		err := modifyFunc(pg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute modification function on resource %s", name)
		}
	}

	// Create or update the resource
	_, err := controllerutil.CreateOrUpdate(c.Ctx, c.Client, pg, func() error {
		pg.Spec.Type = deploymentType
		pg.Spec.Tier = tier
		pg.Spec.SecretRef = &croType.SecretRef{
			Name:      secretName,
			Namespace: secretNs,
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to reconcile postgres request for %s", name)
	}

	return pg, nil
}

// ReconcileRedis creates or updates a redis custom resource
func (c CloudResourceService) ReconcileRedis(productName, deploymentType, tier, name, ns, secretName, secretNs string, modifyFunc modifyResourceFunc) (*v1alpha1.Redis, error) {
	r := &v1alpha1.Redis{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels: map[string]string{
				"productName": productName,
			},
		},
	}

	// execute logic to modify the resource before creation
	// e.g. add owner refs
	if modifyFunc != nil {
		err := modifyFunc(r)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute modification function on resource %s", name)
		}
	}

	// Create or update the resource
	_, err := controllerutil.CreateOrUpdate(c.Ctx, c.Client, r, func() error {
		r.Spec.Type = deploymentType
		r.Spec.Tier = tier
		r.Spec.SecretRef = &croType.SecretRef{
			Name:      secretName,
			Namespace: secretNs,
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to reconcile redis request for %s", name)
	}

	return r, nil
}
