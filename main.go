package main

import (
	"context"
	"database/sql"
	"fmt"
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

	client, err := newPulsarClient(conf.Input)
	if err != nil {
		log.Fatalln("Could not instantiate Pulsar client:", err.Error())
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            conf.Input.Topic,
		SubscriptionName: "test",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	for i := 0; i < 10; i++ {
		// may block here
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
	// _, err = NewDbClient(conf.Output.Rdb)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

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

func newPulsarClient(c config.Pulsar) (pulsar.Client, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               c.Host,
		OperationTimeout:  timeoutSec * time.Second,
		ConnectionTimeout: timeoutSec * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewDbClient(c config.Rdb) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, c.User+":"+c.Password+"@tcp("+c.Host+")/"+c.Schema)
	if err != nil {
		return nil, err
	}
	return db, nil
}
