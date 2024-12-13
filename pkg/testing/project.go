package testing

import (
	"context"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type Project struct {
	internal.Workspace
	afero.Fs
	LoadFunc     func(context.Context, string) (pkg.Workspace, error)
	ParseFunc    func(string) (string, error)
	PathValue    string
	ReadFileFunc func(string) ([]byte, error)
	WalkFunc     func(filepath.WalkFunc) error
}

// Name implements pkg.Project.
func (p *Project) Name() string {
	return "Test"
}

// Load implements pkg.Project.
func (p *Project) Load(ctx context.Context, path string) (pkg.Workspace, error) {
	if p.LoadFunc == nil {
		panic("unimplemented")
	}

	return p.LoadFunc(ctx, path)
}

// Parse implements pkg.Project.
func (p *Project) Parse(path string) (string, error) {
	if p.ParseFunc == nil {
		return p.Workspace.Parse(path)
	}

	return p.ParseFunc(path)
}

// Path implements pkg.Project.
func (p *Project) Path() string {
	return p.PathValue
}

// ReadFile implements pkg.Project.
func (p *Project) ReadFile(filename string) ([]byte, error) {
	if p.ReadFileFunc == nil {
		return p.Workspace.ReadFile(filename)
	}

	return p.ReadFileFunc(filename)
}

// Walk implements pkg.Project.
func (p *Project) Walk(walk filepath.WalkFunc) error {
	if p.WalkFunc == nil {
		return p.Workspace.Walk(walk)
	}

	return p.WalkFunc(walk)
}

var _ pkg.Project = &Project{}
