package project

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/go/vcs/git"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type Git struct {
	internal.Workspace
	loaders []pkg.Loader
}

// Load implements pkg.Project.
func (g *Git) Load(ctx context.Context, path string) (pkg.Workspace, error) {
	path, err := g.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("parsing path %s: %w", path, err)
	}

	stat, err := g.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("reading workspace: %w", err)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("workspace must be a directory: %s", path)
	}

	return internal.Workspace{
		Fs:   g.Fs,
		Root: path,
	}, nil
}

// Path implements pkg.Project.
func (g *Git) Path() string {
	return g.Root
}

// Walk implements pkg.Project.
func (g *Git) Walk(walk filepath.WalkFunc) error {
	return afero.Walk(g.Fs, g.Root, walk)
}

func LocalGit(ctx context.Context) (pkg.Project, error) {
	root, err := git.Root(ctx)
	if err != nil {
		return nil, fmt.Errorf("locating git root: %w", err)
	}

	fs := afero.NewOsFs()
	stat, err := fs.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("loading git workspace: %w", err)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("project must be a directory: %w", err)
	}

	return &Git{
		Workspace: internal.Workspace{
			Root: root,
			Fs:   fs,
		},
		loaders: []pkg.Loader{},
	}, nil
}
