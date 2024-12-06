package config

import (
	"fmt"

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
