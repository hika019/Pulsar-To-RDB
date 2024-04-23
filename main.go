package main

import (
	"fmt"
	"log"

	"github.com/hika019/Pulsar-To-RDB.git/config"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalln(err.Error())
	}
	conf, err := config.LoadConfig(env)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}
