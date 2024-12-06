package workspace

import (
	"errors"
	"io/fs"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
	"github.com/unmango/thecluster/pkg"
	"github.com/unmango/thecluster/pkg/context"
	"github.com/unmango/thecluster/pkg/workspace/pulumi"
)

var (
	loaders        = []pkg.Loader{pulumi.Loader}
	ignorePrefixes = []string{".git", ".vscode", ".make"}
)

func List(ctx pkg.Context) (iter.Seq[pkg.Workspace], error) {
	ws := []pkg.Workspace{}
	err := context.Walk(ctx,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() || path == "" {
				return nil
			}
			for _, p := range ignorePrefixes {
				if strings.HasPrefix(path, p) {
					return nil
				}
			}

			if w, err := Load(ctx, path); err != nil {
				log.Debugf("loading workspace: %s", err)
			} else {
				ws = append(ws, w)
			}

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return slices.Values(ws), nil
}

func Load(ctx pkg.Context, path string) (w pkg.Workspace, err error) {
	errs := []error{}
	for _, l := range loaders {
		if w, err = l.Load(ctx, path); err == nil {
			return
		} else {
			errs = append(errs, err)
		}
	}

	return nil, errors.Join(errs...)
}
