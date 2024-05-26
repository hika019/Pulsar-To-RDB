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
	logLevel := os.Getenv("INFO")
	env := Env{ConfPath: confPath, LogLevel: logLevel}
	return env, nil
}
