package core

import (
	"github.com/nats-io/nats.go"
	"log"
)

func NewMQ() *nats.Conn {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
