package bluetoothfunctions

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
)

func Discoverer(ch chan string) {
	fmt.Printf("Discoverer")
	d, err := gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Println("State:", s)
		switch s {
		case gatt.StatePoweredOn:
			fmt.Println("scanning...")
			d.Scan([]gatt.UUID{}, false)
			return
		default:
			d.StopScanning()
		}
	}

	onPeriphDiscovered := func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		if len(a.Services) > 0 {
			if AvertisedService.Equal(a.Services[0]) {
				//fmt.Printf("\nFound 8Power Peripheral ID: (%s)\n", p.ID())
				ch <- p.ID()
			}
		}
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
	select {}
}
