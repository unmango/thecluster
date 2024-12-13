package config

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"

	"github.com/ianlewis/go-gitignore"
	"github.com/unmango/thecluster/pkg"
	"gopkg.in/yaml.v3"
)

var SupportedNames = []string{
	".thecluster.yml",
	".thecluster.yaml",
	".thecluster",
	"thecluster.yml",
	"thecluster.yaml",
}

func Load(project pkg.Project, path string) (*pkg.Config, error) {
	data, err := project.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config pkg.Config
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	} else {
		return &config, nil
	}
}

func Locate(project pkg.Project) (config string, err error) {
	var ignore gitignore.GitIgnore
	ignore, err = gitignore.NewFromFile(
		filepath.Join(project.Path(), ".gitignore"),
	)
	if err != nil {
		return
	}

	err = project.Walk(
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if ignore.Ignore(path) {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}

			base := filepath.Base(path)
			if !slices.Contains(SupportedNames, base) {
				return nil
			}

			config = path
			return filepath.SkipAll
		},
	)

	return
}
