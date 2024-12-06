package workspace

import (
	"fmt"

	"github.com/unmango/thecluster/pkg"
)

func Install(ctx pkg.Context, work pkg.Workspace) error {
	if i, ok := work.(pkg.Installer); !ok {
		return fmt.Errorf("workspace does not support installing dependencies: %s", work)
	} else {
		return i.Install(ctx)
	}
}
