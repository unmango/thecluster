package context

import (
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/thecluster/pkg"
)

func Walk(ctx pkg.Context, fn filepath.WalkFunc) error {
	return afero.Walk(ctx.Fs(), ctx.Root(), fn)
}
