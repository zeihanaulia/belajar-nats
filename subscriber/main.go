package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	subject := "contoh.subject"
	opts := []nats.Option{nats.Name("Contoh NATS Subsriber")}
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf(" Pesan diterima dari [%s]: '%s'", msg.Subject, string(msg.Data))
	})

	log.Printf("Mendengarkan pada subject [%s]", subject)
	runtime.Goexit()
}
