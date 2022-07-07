package main

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/imokech/distributed-golang-app-with-rabbitmq/src/distributed/datamanager"
	"github.com/imokech/distributed-golang-app-with-rabbitmq/src/distributed/dto"
	"github.com/imokech/distributed-golang-app-with-rabbitmq/src/distributed/qutils"
)

const url = "amqp://guest:guest@localhost:5762"

func main() {
	conn, ch := qutils.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		qutils.PersistReadingsQueue,
		"",    // consumer
		false, // autoAck
		true,  // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)

	if err != nil {
		log.Fatalln("Failed to get access to messages")
	}

	for msg := range msgs {
		buf := bytes.NewReader(msg.Body)
		dec := gob.NewDecoder(buf)
		sd := &dto.SensorMessage{}
		dec.Decode(sd)

		err := datamanager.SaveReading(sd)

		if err != nil {
			log.Printf(
				"Failed to save reading from sensor %v. Error: %s",
				sd.Name,
				err.Error(),
			)
		} else {
			msg.Ack(false)
		}
	}
}
