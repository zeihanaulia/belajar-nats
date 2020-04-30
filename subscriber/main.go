package main

import (
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	subject := "contoh.subject"
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts := []nats.Option{nats.Name("Contoh NATS Subsriber")}
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))

	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf(" Pesan diterima dari [%s]: '%s'", msg.Subject, string(msg.Data))
	})
	nc.Flush()

	log.Printf("Mendengarkan pada subject [%s]", subject)
	runtime.Goexit()
}
