package project

import (
	"context"

	"github.com/unmango/devctl/pkg/work"
)

type Project struct {
	Dir work.Directory
}

func Load(ctx context.Context) (*Project, error) {
	if dir, err := work.Load(ctx); err != nil {
		return nil, err
	} else {
		return &Project{Dir: dir}, nil
	}
}
