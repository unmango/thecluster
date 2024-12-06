package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/afero"
)

type Workspace struct {
	fs   afero.Fs
	root string
}

func (w Workspace) String() string {
	return w.Name()
}

func (w Workspace) Name() string {
	return filepath.Base(w.root)
}

func (w Workspace) Fs() afero.Fs {
	return w.fs
}

func (w Workspace) Path(path string) string {
	return filepath.Join(w.root, path)
}

func (w Workspace) Parse(path string) (string, error) {
	if !filepath.IsAbs(path) {
		return path, nil
	} else {
		return filepath.Rel(w.root, path)
	}
}

func LoadWorkspace(fsys afero.Fs, root string) (w Workspace, err error) {
	var stat fs.FileInfo
	stat, err = fsys.Stat(root)
	if err != nil {
		return w, fmt.Errorf("unable to load workspace: %w", err)
	}
	if !stat.IsDir() {
		return w, fmt.Errorf("workspace must be a directory: %s", root)
	}

	w.fs = afero.NewBasePathFs(fsys, root)
	w.root = root

	return
}
