package bluetoothfunctions

import (
	"github.com/RobHumphris/bluetooth-tests/global"
	"github.com/paypal/gatt"
)

func ListDiscovered(p gatt.Peripheral, a *gatt.Advertisement) {
	LogPeripheralData(p)
	LogAdvertisementData(a)
}

func Discoverer(ch chan gatt.Peripheral) {
	global.Debugf("Discoverer")

	onStateChanged := func(d gatt.Device, s gatt.State) {
		global.Debugf("State: %s", s)
		switch s {
		case gatt.StatePoweredOn:
			global.Debugf("scanning...")
			d.Scan([]gatt.UUID{}, false)
			return
		default:
			d.StopScanning()
		}
	}

	onPeriphDiscovered := func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		if len(a.Services) > 0 {
			if AvertisedServiceUUID.Equal(a.Services[0]) {
				ch <- p
			}
		}
	}

	// Register handlers.
	HCIDevice.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	HCIDevice.Init(onStateChanged)
	select {}
}
