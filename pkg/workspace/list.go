package workspace

import (
	"io/fs"

	"github.com/spf13/afero"
	"github.com/unmango/go/iter"
	"github.com/unmango/go/slices"
	"github.com/unmango/thecluster/pkg"
	"github.com/unmango/thecluster/pkg/context"
	"github.com/unmango/thecluster/pkg/workspace/pulumi"
)

func List(ctx pkg.Context) (iter.Seq[pkg.Workspace], error) {
	ws := []pkg.Workspace{}
	err := context.Walk(ctx,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() || !Exists(ctx.Fs(), path) {
				return nil
			}

			w, err := pulumi.LoadWorkspace(ctx.Fs(), path)
			if err != nil {
				return err
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

func Exists(fs afero.Fs, path string) bool {
	return pulumi.IsWorkspace(fs, path)
}
