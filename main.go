package main

import (
	"log"
	"net/http"

	"github.com/RobHumphris/bluetooth-tests/bluetooth-functions"
	"github.com/RobHumphris/bluetooth-tests/data"
	"github.com/RobHumphris/bluetooth-tests/handlers"
)

var port = ":8000"
var discoveryChan = make(chan string)

func monitorDiscoveries(ch chan string) {
	for {
		select {
		case macAddress := <-discoveryChan:
			data.DiscoveredPeripherals.Store(macAddress)
		}
	}
}

func main() {
	handlers.SetupHandlers()

	go bluetoothfunctions.Discoverer(discoveryChan)
	go monitorDiscoveries(discoveryChan)

	log.Printf("Server listening on port %v\n", port)
	http.ListenAndServe(port, handlers.StatsMiddleware)
	select {}
}
