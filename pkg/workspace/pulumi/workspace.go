package pulumi

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/thecluster/pkg"
)

type workspace struct {
	fs   afero.Fs
	root string
}

// Fs implements pkg.Workspace.
func (w *workspace) Fs() afero.Fs {
	return w.fs
}

// Name implements pkg.Workspace.
func (w *workspace) Name() string {
	return filepath.Base(w.root)
}

func IsWorkspace(fs afero.Fs, path string) bool {
	stat, err := fs.Stat(filepath.Join(path, "Pulumi.yaml"))

	return err == nil && !stat.IsDir()
}

func LoadWorkspace(fs afero.Fs, path string) (pkg.Workspace, error) {
	if !IsWorkspace(fs, path) {
		return nil, fmt.Errorf("not a Pulumi workspace: %s", path)
	}

	return &workspace{
		fs:   afero.NewBasePathFs(fs, path),
		root: path,
	}, nil
}
