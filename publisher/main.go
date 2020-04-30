package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	subject := "contoh.subject"
	msg := "hello abc 123"

	opts := []nats.Option{nats.Name("Contoh NATS Publisher")}
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	err = nc.Publish(subject, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Diterbitkan ke [%s] : '%s'\n", subject, msg)
	}
}
