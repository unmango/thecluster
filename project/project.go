package project

import (
	"context"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"

	"github.com/unmango/go/iter"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/work"
)

var blacklist = []string{
	"node_modules",
	".git",
	".make",
}

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

func (p *Project) Workspaces(ctx context.Context) (ws iter.Seq[auto.Workspace], err error) {
	err = afero.Walk(p.Dir.Fs(), "",
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			blacklisted := func(s string) bool {
				return strings.Contains(path, s)
			}
			if info.IsDir() && slices.ContainsFunc(blacklist, blacklisted) {
				return filepath.SkipDir
			}

			if p.IsWorkspace(path, info) {
				if work, err := auto.NewLocalWorkspace(ctx); err != nil {
					return err
				} else {
					ws = iter.Append(ws, work)
				}
			}

			return nil
		},
	)
	return
}

func (p *Project) IsWorkspace(path string, info fs.FileInfo) bool {
	if !info.IsDir() {
		return false
	}

	pyaml := filepath.Join(path, "Pulumi.yaml")
	ok, _ := afero.Exists(p.Dir.Fs(), pyaml)
	return ok
}
