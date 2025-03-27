package core

import (
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
)

func NewMQ() *nats.Conn {
	conn, err := nats.Connect(viper.GetString("nats.address"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
