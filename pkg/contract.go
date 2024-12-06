package pkg

import (
	"context"

	"github.com/spf13/afero"
)

type Context interface {
	context.Context
	Workspace
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
	Install(Context) error
}
