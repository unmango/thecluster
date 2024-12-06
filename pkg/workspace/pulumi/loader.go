package pulumi

import (
	"github.com/unmango/thecluster/pkg"
)

type loader string

var Loader loader = "Pulumi"

func (loader) Load(ctx pkg.Context, path string) (pkg.Workspace, error) {
	return Load(ctx, path)
}
