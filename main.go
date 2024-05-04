package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/hika019/Pulsar-To-RDB.git/config"

	_ "github.com/go-sql-driver/mysql"
)

const timeoutSec = 30

func main() {
	conf, err := initConf()
	if err != nil {
		log.Fatal("Could not init Config:", err.Error())
	}

	//consume(conf.Input)
	_, err = NewDbClient(conf.Output.Rdb)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func initConf() (config.Config, error) {
	env, err := config.LoadEnv()
	if err != nil {
		return config.Config{}, err
	}
	conf, err := config.LoadConfig(env)

	if err != nil {
		return config.Config{}, err
	}
	return conf, nil
}

func consume(c config.Pulsar) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               c.Host,
		OperationTimeout:  timeoutSec * time.Second,
		ConnectionTimeout: timeoutSec * time.Second,
	})
	if err != nil {
		log.Fatalln("Could not instantiate Pulsar client:", err.Error())
	}

	defer client.Close()
}

func NewDbClient(c config.Rdb) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, c.User+":"+c.Password+"@tcp("+c.Host+")/"+c.Schema)
	if err != nil {
		return nil, err
	}
	return db, nil
}
