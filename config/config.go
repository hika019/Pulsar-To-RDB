package config

import (
	"encoding/json"
	"log"
	"os"
)

func LoadConfig(env Env) (Config, error) {
	f, err := os.Open(env.ConfPath)
	if err != nil {
		return Config{}, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	var conf Config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}
