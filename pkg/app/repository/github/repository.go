package github

import (
	"context"
	"iter"
	"net/http"
	"slices"

	"github.com/unmango/aferox/github"
	theclusterv1alpha1 "github.com/unmango/thecluster/gen/dev/unmango/thecluster/v1alpha1"
)

type Repository struct {
	name, owner string
}

// Name implements app.Repository.
func (g *Repository) Name() string {
	return g.name
}

// ListApps implements app.Repository.
func (g *Repository) ListApps(ctx context.Context) (iter.Seq[*theclusterv1alpha1.App], error) {
	fs := github.NewFs(github.NewClient(http.DefaultClient))
	appDir, err := fs.Open("")
	if err != nil {
		return nil, err
	}

	dirs, err := appDir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	apps := []*theclusterv1alpha1.App{}
	for _, d := range dirs {
		apps = append(apps, &theclusterv1alpha1.App{
			Name: &d,
		})
	}

	return slices.Values(apps), nil
}
