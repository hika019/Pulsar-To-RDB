package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/hika019/Pulsar-To-RDB.git/config"

	_ "github.com/go-sql-driver/mysql"
)

const timeoutSec = 30

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	conf, env, err := initConf()
	if err != nil {
		slog.Error("Could not init Config", err)
		os.Exit(1)
	}

	logLevel := slog.LevelInfo
	switch env.LogLevel {
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "ERROR":
		logLevel = slog.LevelError
	}
	slog.SetLogLoggerLevel(logLevel)

	client, err := newPulsarClient(conf.Input)
	if err != nil {
		slog.Error("Could not instantiate Pulsar client", err)
		os.Exit(1)
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            conf.Input.Topic,
		SubscriptionName: "test",
		Type:             pulsar.Shared,
	})

	if err != nil {
		slog.Error("Could not create consumer", err)
		os.Exit(1)
	}

	defer func() {
		if err := consumer.Unsubscribe(); err != nil {
			slog.Error("Unsubscribe failed", err)
			os.Exit(1)
		}
		slog.Info("Unsubscribe is Succeed")
		consumer.Close()
		slog.Info("Close consumer")
	}()

	db, err := NewDbClient(conf.Output.Rdb)
	if err != nil {
		slog.Error("Could not instantiate DB client")
		os.Exit(1)
	}
	defer db.Close()

	d := filepath.Dir(conf.Output.File.Path)
	if err := os.MkdirAll(d, 0755); err != nil {
		slog.Error("Could not create dir", err)
		os.Exit(1)
	}
	f, err := os.OpenFile(conf.Output.File.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		slog.Error("Could not open file", err)
		os.Exit(1)
	}
	defer f.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			slog.Warn(err.Error())
		}

		payload := make(map[string]string)
		err = json.Unmarshal(msg.Payload(), &payload)
		if err != nil {
			slog.Warn(err.Error())
			continue
		}

		var m map[string]any
		json.Unmarshal([]byte(payload["message"]), &m)

		ins, err := db.Prepare(conf.Output.Rdb.Statement[0])
		if err != nil {
			slog.Warn(err.Error())
		}

		values := []any{}
		for i, s := range conf.Output.Rdb.Statement {
			if i == 0 {
				continue
			}
			values = append(values, m[s])
		}

		bytes, err := json.Marshal(m)
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			continue
		}

		fmt.Fprintln(f, string(bytes))

		_, err = ins.Exec(values...)
		if err != nil {
			slog.Warn(err.Error())
			consumer.Nack(msg)
			continue
		}
		consumer.Ack(msg)
	}
}

func initConf() (config.Config, config.Env, error) {
	env, err := config.LoadEnv()
	if err != nil {
		return config.Config{}, config.Env{}, err
	}
	conf, err := config.LoadConfig(env)

	if err != nil {
		return config.Config{}, config.Env{}, err
	}
	return conf, env, nil
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
