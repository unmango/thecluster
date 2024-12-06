package pulumi

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/spf13/afero"
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/subject"
	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

type Workspace struct {
	internal.Workspace
	events rx.Subject[pkg.ProgressEvent]
	pulumi auto.Workspace
}

func (w *Workspace) Install(ctx context.Context) error {
	return w.pulumi.Install(ctx, &auto.InstallOptions{
		// TODO
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
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

func Load(ctx pkg.Context, path string) (*Workspace, error) {
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

	ws, err := auto.NewLocalWorkspace(ctx,
		auto.WorkDir(ctx.Path(rel)),
	)
	if err != nil {
		return nil, fmt.Errorf("loading workspace: %w", err)
	}

	return &Workspace{
		Workspace: w,
		pulumi:    ws,
		events:    subject.New[pkg.ProgressEvent](),
	}, nil
}
