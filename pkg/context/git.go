package context

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
	"github.com/unmango/go/vcs/git"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type repo struct {
	context.Context
	internal.Workspace
}

func LocalRepo(ctx context.Context) (pkg.Context, error) {
	root, err := git.Root(ctx)
	if err != nil {
		return nil, fmt.Errorf("locating git root: %w", err)
	}

	ws, err := internal.LoadWorkspace(afero.NewOsFs(), root)
	if err != nil {
		return nil, fmt.Errorf("loading workspace: %w", err)
	}

	return &repo{ctx, ws}, nil
}
