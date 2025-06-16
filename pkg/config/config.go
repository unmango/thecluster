package config

import (
	"github.com/spf13/viper"
	"github.com/unmango/thecluster/pkg/contract"
)

const (
	TargetRepositoryKey = "git_repo"
)

var (
	global = New()
)

func Global() contract.Config {
	return global
}

type Config struct {
	viper *viper.Viper
}

func New() contract.Config {
	v := viper.NewWithOptions()
	v.SetEnvPrefix("THECLUSTER")
	v.AutomaticEnv()

	return &Config{v}
}

// TargetRepository implements contract.Config.
func (c *Config) TargetRepository() string {
	return c.viper.GetString(TargetRepositoryKey)
}
