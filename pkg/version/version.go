package version

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/unmango/thecluster/pkg"
)

const DefaultPath = ".versions"

func Get(project pkg.Project, name string) (string, error) {
	if err := ensure(project); err != nil {
		return "", err
	}

	data, err := project.ReadFile(filepath.Join(DefaultPath, name))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func ensure(project pkg.Project) error {
	if stat, err := project.Stat(DefaultPath); err == nil && stat.IsDir() {
		return nil
	} else {
		return project.Mkdir(DefaultPath, os.ModeDir)
	}
}
