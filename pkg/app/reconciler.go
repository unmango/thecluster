package app

import (
	"context"

	"github.com/unmango/thecluster/pkg/contract"
	"github.com/unmango/thecluster/pkg/result"
)

type Reconciler struct {
	cfg contract.Config
}

func FromConfig(cfg contract.Config) contract.App {
	return &Reconciler{cfg}
}

// Reconcile implements contract.App.
func (a *Reconciler) Reconcile(ctx context.Context, req contract.Request) (contract.Result, error) {
	return result.Done(), nil
}
