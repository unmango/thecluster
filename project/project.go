package project

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/iter"
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

func LoadFrom(path string) (*Project, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	} else {
		return &Project{Dir: work.Directory(path)}, nil
	}
}

func (p *Project) Workspaces() (iter.Seq[Workspace], error) {
	ws := iter.Empty[Workspace]()
	err := p.walk(func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ignore, err := p.ignore(path, info); ignore {
			return err
		}
		if _, err := workspace.DetectProjectPathFrom(path); err == nil {
			ws = iter.Append(ws, Workspace(path))
		}

		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return ws, nil
	}
}

func (p *Project) ignore(path string, info fs.FileInfo) (bool, error) {
	if path == "" || !info.IsDir() {
		return true, nil
	}

	blacklisted := func(s string) bool {
		return strings.Contains(path, s)
	}
	if slices.ContainsFunc(blacklist, blacklisted) {
		return true, fs.SkipDir
	}

	return false, nil
}

func (p *Project) walk(walkFn filepath.WalkFunc) error {
	return afero.Walk(p.Dir.Fs(), "",
		func(path string, info fs.FileInfo, err error) error {
			return walkFn(filepath.Join(p.Dir.Path(), path), info, err)
		},
	)
}

// func (p *Project) Workspaces(ctx context.Context) (ws iter.Seq[auto.Workspace], err error) {
// 	err = afero.Walk(p.Dir.Fs(), "",
// 		func(path string, info fs.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}

// 			blacklisted := func(s string) bool {
// 				return strings.Contains(path, s)
// 			}
// 			if info.IsDir() && slices.ContainsFunc(blacklist, blacklisted) {
// 				return filepath.SkipDir
// 			}

// 			if p.IsWorkspace(path, info) {
// 				if work, err := auto.NewLocalWorkspace(ctx); err != nil {
// 					return err
// 				} else {
// 					ws = iter.Append(ws, work)
// 				}
// 			}

// 			return nil
// 		},
// 	)
// 	return
// }

// func (p *Project) IsWorkspace(path string, info fs.FileInfo) bool {
// 	if !info.IsDir() {
// 		return false
// 	}

// 	pyaml := filepath.Join(path, "Pulumi.yaml")
// 	ok, _ := afero.Exists(p.Dir.Fs(), pyaml)
// 	return ok
// }
