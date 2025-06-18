package app

import (
	"context"
	"iter"

	theclusterv1alpha1 "github.com/unmango/thecluster/gen/dev/unmango/thecluster/v1alpha1"
)

type Repository interface {
	Name() string
	ListApps(context.Context) (iter.Seq[*theclusterv1alpha1.App], error)
}
