package pulumi

import (
	"fmt"

	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type loader string

var Loader loader = "Pulumi"

func (loader) Load(ctx pkg.Context, path string) (pkg.Workspace, error) {
	rel, err := ctx.Parse(path)
	if err != nil {
		return nil, err
	}

	if !IsWorkspace(ctx.Fs(), rel) {
		return nil, fmt.Errorf("not a pulumi workspace: %s", path)
	}

	w, err := internal.LoadWorkspace(ctx.Fs(), rel)
	if err != nil {
		return nil, fmt.Errorf("loading workspace: %w", err)
	}

	return &Workspace{w}, nil
}
