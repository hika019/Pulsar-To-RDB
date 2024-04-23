package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Env{}, err
	}

	confPath := os.Getenv("CONFIG_PATH")
	env := Env{ConfPath: confPath}
	return env, nil
}
