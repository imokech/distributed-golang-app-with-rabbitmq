package main

import (
	"fmt"

	"github.com/imokech/distributed-golang-app-with-rabbitmq/src/distributed/coordinator"
)

var dc *coordinator.DatabseConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseConsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
