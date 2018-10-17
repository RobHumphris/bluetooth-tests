package main

import (
	"net/http"

	bluetoothfunctions "github.com/RobHumphris/bluetooth-tests/bluetooth-functions"
	"github.com/RobHumphris/bluetooth-tests/data"
	"github.com/RobHumphris/bluetooth-tests/global"
	"github.com/RobHumphris/bluetooth-tests/http-handlers"
	"github.com/paypal/gatt"
)

var port = ":8000"
var discoveryChan = make(chan gatt.Peripheral)

func monitorDiscoveries(ch chan gatt.Peripheral) {
	for {
		select {
		case peripheral := <-discoveryChan:
			data.DiscoveredPeripherals.Store(peripheral)
		}
	}
}

func main() {
	httphandlers.SetupHandlers()

	go bluetoothfunctions.Discoverer(discoveryChan)
	go monitorDiscoveries(discoveryChan)

	global.Debugf("Server listening on port %v\n", port)
	http.ListenAndServe(port, httphandlers.StatsMiddleware)
	select {}
}
