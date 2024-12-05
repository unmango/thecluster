package pkg

import (
	"context"

	"github.com/spf13/afero"
)

type Context interface {
	context.Context
	Root() string
	Fs() afero.Fs
}

type Named interface {
	Name() string
}

type Workspace interface {
	Named
	Fs() afero.Fs
}
