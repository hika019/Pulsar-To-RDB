package main

import (
	"encoding/json"
	"log"
	"os"
)

const config = "./config.json"

func readConfig() (Config, error) {
	f, err := os.Open(config)
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
