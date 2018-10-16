package bluetoothfunctions

import (
	"log"

	"github.com/paypal/gatt"
)

const (
	AdvertisedServiceString     = "88881310deadbea71523785feab7e123"
	CommandCharacteristicString = "88881311deadbea71523785feab7e123"
	DataCharacteristicString    = "88881312deadbea71523785feab7e123"
)

var AvertisedService gatt.UUID
var CommandCharacteristic gatt.UUID
var DataCharacteristic gatt.UUID

var DefaultClientOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
}

func init() {
	var err error
	AvertisedService, err = gatt.ParseUUID(AdvertisedServiceString)
	if err != nil {
		log.Fatalf("error parsing AdvertisedServiceString: %v", err)
	}

	CommandCharacteristic, err = gatt.ParseUUID(CommandCharacteristicString)
	if err != nil {
		log.Fatalf("error parsing CommandCharacteristicString: %v", err)
	}

	DataCharacteristic, err = gatt.ParseUUID(DataCharacteristicString)
	if err != nil {
		log.Fatalf("error parsing DataCharacteristicString: %v", err)
	}
}
