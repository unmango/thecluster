package testing

import (
	"context"

	"github.com/onsi/ginkgo/v2"
	"github.com/spf13/afero"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type Context struct {
	context.Context
	FsValue   afero.Fs
	NameValue string
	ParseFunc func(string) (string, error)
}

// Name implements pkg.Context.
func (ctx *Context) Name() string {
	return ctx.NameValue
}

// Parse implements pkg.Context.
func (ctx *Context) Parse(p string) (string, error) {
	if ctx.ParseFunc == nil {
		panic("unimplemented")
	}

	return ctx.ParseFunc(p)
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
	fs := afero.NewBasePathFs(afero.NewOsFs(), tmp)

	return &Context{
		Context:   context.TODO(),
		FsValue:   fs,
		NameValue: tmp,
	}
}

func DefaultContext(ctx context.Context, fs afero.Fs) *Context {
	return &Context{
		Context:   ctx,
		FsValue:   fs,
		NameValue: "Default",
		ParseFunc: func(s string) (string, error) {
			return internal.Workspace{}.Parse(s)
		},
	}
}
