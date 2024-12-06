package testing

import (
	"context"
	"path/filepath"

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
	PathFunc  func(string) string
}

// Path implements pkg.Context.
func (ctx *Context) Path(s string) string {
	if ctx.PathFunc == nil {
		panic("unimplemented")
	}

	return ctx.PathFunc(s)
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
		ParseFunc: func(s string) (string, error) {
			return filepath.Rel(tmp, s)
		},
		PathFunc: func(s string) string {
			return filepath.Join(tmp, s)
		},
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
		PathFunc: func(s string) string {
			return internal.Workspace{}.Path(s)
		},
	}
}
