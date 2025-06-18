package repository

import (
	"context"

	theclusterv1alpha1 "github.com/unmango/thecluster/gen/dev/unmango/thecluster/v1alpha1"
	"github.com/unmango/thecluster/pkg/app"
)

func Get(ctx context.Context, app *theclusterv1alpha1.App) (app.Repository, error) {
	return &GitHub{name: *app.Repository}, nil
}
