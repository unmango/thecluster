package pulumi

import (
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/thecluster/internal"
)

type Workspace struct {
	internal.Workspace
}

func IsWorkspace(fs afero.Fs, path string) bool {
	stat, err := fs.Stat(
		filepath.Join(path, "Pulumi.yaml"),
	)

	return err == nil && !stat.IsDir()
}
