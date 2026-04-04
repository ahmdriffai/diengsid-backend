package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./config/")
	cfg.AddConfigPath("./")
	cfg.AddConfigPath("../config/")
	cfg.AddConfigPath("../")
	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	setDefaults(cfg)

	if err := cfg.ReadInConfig(); err != nil {
		if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
			cfg.SetConfigFile(configPath)
			if readErr := cfg.ReadInConfig(); readErr == nil {
				return cfg
			}
		}
	}

	return cfg
}

func setDefaults(cfg *viper.Viper) {
	cfg.SetDefault("app.name", "diengid-backend")
	cfg.SetDefault("app.prefork", false)
	cfg.SetDefault("web.port", 8080)
	cfg.SetDefault("log.level", 4)
}
