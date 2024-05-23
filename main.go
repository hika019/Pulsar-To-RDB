package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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

	defer func() {
		if err := consumer.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Unsubscribe")
		consumer.Close()
		fmt.Println("Close consumer")
	}()

	db, err := NewDbClient(conf.Output.Rdb)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		payload := make(map[string]string)
		err = json.Unmarshal(msg.Payload(), &payload)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		fmt.Println()
		fmt.Println(payload["message"])
		fmt.Println(reflect.TypeOf(payload["message"]))

		var m map[string]any
		json.Unmarshal([]byte(payload["message"]), &m)

		fmt.Println("aaa")
		fmt.Println(m)

		ins, err := db.Prepare(conf.Output.Rdb.Statement[0])
		if err != nil {
			log.Fatalln(err)
		}

		values := []any{}
		for i, s := range conf.Output.Rdb.Statement {
			if i == 0 {
				continue
			}
			values = append(values, m[s])

		}
		_, err = ins.Exec(values...)

		if err != nil {
			log.Fatal(err)
			consumer.Nack(msg)
			continue
		}
		consumer.Ack(msg)
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
