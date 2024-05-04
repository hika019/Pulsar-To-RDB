package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(env Env) (Config, error) {
	b, err := os.ReadFile(env.ConfPath)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	yaml.Unmarshal(b, &conf)

	return conf, nil
}
