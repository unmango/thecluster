package config

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"

	"github.com/ianlewis/go-gitignore"
	"github.com/spf13/afero"
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

func Load(ctx pkg.Context, path string) (*pkg.Config, error) {
	data, err := afero.ReadFile(ctx.Fs(), path)
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

func Locate(ctx pkg.Context) (config string, err error) {
	var ignore gitignore.GitIgnore
	ignore, err = gitignore.NewFromFile(
		ctx.Path(".gitignore"),
	)
	if err != nil {
		return
	}

	err = afero.Walk(ctx.Fs(), "",
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
