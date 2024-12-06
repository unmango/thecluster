package version

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/thecluster/pkg"
)

const DefaultPath = ".versions"

func Get(ctx pkg.Context, name string) (string, error) {
	if err := ensure(ctx.Fs()); err != nil {
		return "", err
	}

	data, err := afero.ReadFile(ctx.Fs(), filepath.Join(DefaultPath, name))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func ensure(fs afero.Fs) error {
	if stat, err := fs.Stat(DefaultPath); err == nil && stat.IsDir() {
		return nil
	} else {
		return fs.Mkdir(DefaultPath, os.ModeDir)
	}
}
