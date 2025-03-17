package pkg

import (
	"context"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/go/rx"
)

type Named interface {
	Name() string
}

type Workspace interface {
	afero.Fs
	Named
	Parse(string) (string, error)
	ReadFile(string) ([]byte, error)
	Walk(filepath.WalkFunc) error
}

type Project interface {
	Workspace
	Load(context.Context, string) (Workspace, error)
	Path() string
}

type Loader interface {
	Load(context.Context, Project, string) (Workspace, error)
}

type Installer interface {
	Install(context.Context) error
}

type ProgressEvent struct {
	Message string
}

type Observable interface {
	rx.Observable[ProgressEvent]
}

type Config struct {
	Deps map[string][]string `yaml:"deps"`
}
