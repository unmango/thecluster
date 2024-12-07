package config

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/ianlewis/go-gitignore"
	"github.com/spf13/afero"
	"github.com/unmango/thecluster/pkg"
	"gopkg.in/yaml.v3"
)

func Load(ctx pkg.Context) (*pkg.Config, error) {
	data, err := afero.ReadFile(ctx.Fs(), ".thecluster.yml")
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

func Locate(ctx pkg.Context) (path string, err error) {
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

			return nil
		},
	)

	return
}
