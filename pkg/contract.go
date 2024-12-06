package pkg

import (
	"context"

	"github.com/spf13/afero"
	"github.com/unmango/go/rx"
)

type Context interface {
	context.Context
	Workspace
	Path(string) string
}

type Named interface {
	Name() string
}

type Workspace interface {
	Named
	Fs() afero.Fs
	Parse(string) (string, error)
}

type Loader interface {
	Load(Context, string) (Workspace, error)
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
