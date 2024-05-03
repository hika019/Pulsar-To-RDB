package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(env Env) (Config, error) {
	fmt.Println(env.ConfPath)
	b, err := os.ReadFile(env.ConfPath)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	yaml.Unmarshal(b, &conf)

	return conf, nil
}
