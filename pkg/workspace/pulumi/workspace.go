package pulumi

import (
	"context"
	"errors"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/unmango/go/rx"
	"github.com/unmango/go/rx/observable"
	"github.com/unmango/go/rx/observer"
	"github.com/unmango/go/rx/subject"
	"github.com/unmango/thecluster/pkg"
)

var ProjectFiles = []string{
	"Pulumi.yaml",
	"Pulumi.yml",
}

type loader string

var Loader loader = "Pulumi"

func (loader) Load(ctx context.Context, project pkg.Project, path string) (pkg.Workspace, error) {
	return Load(ctx, project, path)
}

type Workspace struct {
	pkg.Workspace
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

func IsWorkspace(work pkg.Workspace) bool {
	for _, name := range ProjectFiles {
		path, err := work.Parse(name)
		if err != nil {
			continue
		}

		if stat, err := work.Stat(path); err == nil {
			return !stat.IsDir()
		}
	}

	return false
}

func Load(ctx context.Context, project pkg.Project, path string) (*Workspace, error) {
	work, err := project.Load(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("loading workspace: %w", err)
	}
	if !IsWorkspace(work) {
		return nil, fmt.Errorf("not a pulumi workspace: %s", path)
	}

	pulumi, err := auto.NewLocalWorkspace(ctx,
		auto.WorkDir(path),
	)
	if err != nil {
		return nil, fmt.Errorf("loading workspace: %w", err)
	}

	return &Workspace{
		Workspace: work,
		pulumi:    pulumi,
		events:    subject.New[pkg.ProgressEvent](),
	}, nil
}
