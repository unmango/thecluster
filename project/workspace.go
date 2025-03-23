package project

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

type Workspace string

func (w Workspace) Path() string {
	return w.String()
}

func (w Workspace) String() string {
	return string(w)
}

func (w Workspace) Load(ctx context.Context) (auto.Workspace, error) {
	return auto.NewLocalWorkspace(ctx,
		auto.WorkDir(w.Path()),
	)
}
