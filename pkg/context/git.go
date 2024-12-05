package context

import (
	"context"

	"github.com/spf13/afero"
	"github.com/unmango/go/vcs/git"
	"github.com/unmango/thecluster/pkg"
)

type repo struct {
	context.Context
	root string
	fs   afero.Fs
}

func (ctx *repo) Fs() afero.Fs {
	return ctx.fs
}

func (ctx *repo) Root() string {
	return ctx.root
}

func LocalRepo(ctx context.Context) (pkg.Context, error) {
	if root, err := git.Root(ctx); err != nil {
		return nil, err
	} else {
		fs := afero.NewBasePathFs(afero.NewOsFs(), root)
		return &repo{ctx, root, fs}, nil
	}
}
