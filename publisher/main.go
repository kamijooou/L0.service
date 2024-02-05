package main

import (
	"log"
	"os"
	"publisher/msg"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	var (
		clusterID, clientID, channelName string = "test-cluster", "l0-pub", "l0-channel"
		URL                              string = stan.DefaultNatsURL
	)

	args := os.Args[1:]

	// Connect Options.
	opts := []nats.Option{nats.Name("L0 STAN publisher")}

	// Connect to NATS
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	choice := ``
	switch args[0] {
	case "0":
		choice = msg.Message0
	case "1":
		choice = msg.Message1
	case "2":
		choice = msg.Message2
	case "3":
		choice = msg.Message3
	case "4":
		choice = msg.Message4
	case "5":
		choice = msg.Message5
	case "6":
		choice = msg.Message6
	case "7":
		choice = msg.Message7
	case "8":
		choice = msg.Message8
	case "9":
		choice = msg.Message9
	case "10":
		choice = msg.Message10
	default:
		choice = "{}"
	}

	msg := []byte(choice)

	err = sc.Publish(channelName, msg)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s]: message %s", channelName, args[0])
}
