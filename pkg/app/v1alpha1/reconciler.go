package appv1alpha1

import (
	"context"

	"github.com/unmango/thecluster/pkg/contract"
	"github.com/unmango/thecluster/pkg/result"
)

type Reconciler struct {
	repo string
}

func FromConfig(cfg contract.Config) contract.Reconciler {
	return &Reconciler{
		repo: cfg.TargetRepository(),
	}
}

// Reconcile implements contract.App.
func (a *Reconciler) Reconcile(ctx context.Context, req contract.Request) (contract.Result, error) {
	return result.Done(), nil
}
