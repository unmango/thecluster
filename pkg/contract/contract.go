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

type Resource interface {
	Reconcile(context.Context, Request) (Result, error)
}

type App interface {
	Resource
}

type Config interface {
	TargetRepository() string
}
