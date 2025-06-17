package contract

import (
	"context"
	"time"
)

type Client[T any] interface {
	Get(context.Context, Request) (T, error)
}

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
