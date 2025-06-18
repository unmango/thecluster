package repository

import (
	"context"
	"iter"
	"net/http"

	"github.com/google/go-github/v72/github"
	theclusterv1alpha1 "github.com/unmango/thecluster/gen/dev/unmango/thecluster/v1alpha1"
)

type GitHub struct {
	name string
}

// Name implements app.Repository.
func (g *GitHub) Name() string {
	return g.name
}

// ListApps implements app.Repository.
func (g *GitHub) ListApps(ctx context.Context) (iter.Seq[*theclusterv1alpha1.App], error) {
	client := github.NewClient(http.DefaultClient)
	client.Repositories.GetContents(ctx, "", "", "apps", nil)

	panic("unimplemented")
}
