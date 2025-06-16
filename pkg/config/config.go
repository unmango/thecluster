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

type config struct {
	viper *viper.Viper
}

func New() contract.Config {
	v := viper.NewWithOptions()
	v.SetEnvPrefix("THECLUSTER")
	v.AutomaticEnv()

	return &config{v}
}

// TargetRepository implements contract.Config.
func (c *config) TargetRepository() string {
	return c.viper.GetString(TargetRepositoryKey)
}
