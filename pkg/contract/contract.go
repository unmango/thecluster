package contract

import (
	"context"
	"time"
)

type Request interface {
	Name() string
}

type Result interface {
	RequeueAfter() time.Time
}

type Reconciler interface {
	Reconcile(context.Context, Request) (Result, error)
}

type Config interface {
	TargetRepository() string
}
