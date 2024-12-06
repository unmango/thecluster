package testing

import (
	"context"

	"github.com/onsi/ginkgo/v2"
	"github.com/spf13/afero"
	"github.com/unmango/thecluster/pkg"
)

type Context struct {
	context.Context
	FsValue   afero.Fs
	RootValue string
}

// Name implements pkg.Context.
func (ctx *Context) Name() string {
	panic("unimplemented")
}

// Parse implements pkg.Context.
func (ctx *Context) Parse(string) (string, error) {
	panic("unimplemented")
}

// Fs implements pkg.Context.
func (ctx *Context) Fs() afero.Fs {
	if ctx.FsValue == nil {
		panic("unimplemented")
	}

	return ctx.FsValue
}

var _ pkg.Context = &Context{}

func TempDirContext(t ginkgo.GinkgoTInterface) *Context {
	tmp := t.TempDir()
	fs := afero.NewBasePathFs(
		afero.NewOsFs(),
		tmp,
	)

	return &Context{
		Context:   context.TODO(),
		FsValue:   fs,
		RootValue: tmp,
	}
}
