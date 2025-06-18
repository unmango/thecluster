package appv1alpha1

import (
	"context"

	theclusterv1alpha1 "github.com/unmango/thecluster/gen/dev/unmango/thecluster/v1alpha1"
	"github.com/unmango/thecluster/pkg/app/repository"
	"github.com/unmango/thecluster/pkg/contract"
	"github.com/unmango/thecluster/pkg/reconcile"
	"github.com/unmango/thecluster/pkg/result"
)

type Reconciler struct {
	contract.Client[*theclusterv1alpha1.App]
	repo string
}

func FromConfig(cfg contract.Config) reconcile.Reconciler {
	return &Reconciler{
		repo: cfg.TargetRepository(),
	}
}

// Reconcile implements contract.App.
func (a *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	app, err := a.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	repo, err := repository.Get(ctx, app)
	if err != nil {
		return nil, err
	}

	_, err = repo.ListApps(ctx)
	if err != nil {
		return nil, err
	}

	return result.Done(), nil
}
