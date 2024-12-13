package internal

import (
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

type Workspace struct {
	afero.Fs
	Root string
}

func (w Workspace) Name() string {
	return filepath.Base(w.Root)
}

func (w Workspace) ReadFile(filename string) ([]byte, error) {
	return afero.ReadFile(w, filename)
}

func (w Workspace) String() string {
	return w.Name()
}

func (w Workspace) Parse(path string) (string, error) {
	if strings.HasPrefix(path, w.Root) {
		return path, nil
	} else {
		return filepath.Join(w.Root, path), nil
	}
}

func (w Workspace) Walk(walk filepath.WalkFunc) error {
	return afero.Walk(w.Fs, w.Root, walk)
}
