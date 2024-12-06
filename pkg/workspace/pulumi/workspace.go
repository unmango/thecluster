package pulumi

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/spf13/afero"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type Workspace struct {
	internal.Workspace
}

func (w *Workspace) Install(ctx pkg.Context) error {
	ws, err := auto.NewLocalWorkspace(ctx)
	if err != nil {
		return fmt.Errorf("loading workspace: %w", err)
	}

	return ws.Install(ctx, nil)
}

func IsWorkspace(fs afero.Fs, path string) bool {
	stat, err := fs.Stat(
		filepath.Join(path, "Pulumi.yaml"),
	)
	if err != nil {
		stat, err = fs.Stat(
			filepath.Join(path, "Pulumi.yml"),
		)
	}

	return err == nil && !stat.IsDir()
}

func Load(ctx pkg.Context, path string) (pkg.Workspace, error) {
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
