package pulumi

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/spf13/afero"
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/observable"
	"github.com/unmango/go/rx/observer"
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
	stdout := observable.NewWriter()
	stderr := observable.NewWriter()

	sa := stdout.Subscribe(observer.Anonymous[[]byte]{
		Next: func(b []byte) {
			w.events.OnNext(pkg.ProgressEvent{
				Message: string(b),
			})
		},
	})
	sb := stderr.Subscribe(observer.Anonymous[[]byte]{
		Next: func(b []byte) {
			w.events.OnError(errors.New(string(b)))
		},
	})

	defer sa()
	defer sb()

	return w.pulumi.Install(ctx, &auto.InstallOptions{
		Stdout: stdout,
		Stderr: stderr,
	})
}

func (w *Workspace) Subscribe(obs rx.Observer[pkg.ProgressEvent]) rx.Subscription {
	return w.events.Subscribe(obs)
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
