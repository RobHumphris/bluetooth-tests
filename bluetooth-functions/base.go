package bluetoothfunctions

import (
	"log"

	"github.com/RobHumphris/bluetooth-tests/global"
	"github.com/paypal/gatt"
)

const (
	GenericAccessServiceString              = "1800"
	DeviceNameString                        = "2a00"
	AppearanceString                        = "2a01"
	PeripheralPreferredConnectionParameters = "2a04"

	AdvertisedServiceString     = "88881310deadbea71523785feab7e123"
	CommandCharacteristicString = "88881311deadbea71523785feab7e123"
	DataCharacteristicString    = "88881312deadbea71523785feab7e123"
)

var GenericAccessServiceUUID gatt.UUID
var DeviceNameUUID gatt.UUID

var AvertisedServiceUUID gatt.UUID
var CommandCharacteristicUUID gatt.UUID
var DataCharacteristicUUID gatt.UUID
var HCIDevice gatt.Device

var DefaultClientOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
}

func init() {
	var err error

	GenericAccessServiceUUID, err = gatt.ParseUUID(GenericAccessServiceString)
	if err != nil {
		log.Fatalf("error parsing GenericAccessServiceString: %v", err)
		return
	}

	DeviceNameUUID, err = gatt.ParseUUID(DeviceNameString)
	if err != nil {
		log.Fatalf("error parsing DeviceNameString: %v", err)
		return
	}

	AvertisedServiceUUID, err = gatt.ParseUUID(AdvertisedServiceString)
	if err != nil {
		log.Fatalf("error parsing AdvertisedServiceString: %v", err)
		return
	}

	CommandCharacteristicUUID, err = gatt.ParseUUID(CommandCharacteristicString)
	if err != nil {
		log.Fatalf("error parsing CommandCharacteristicString: %v", err)
		return
	}

	DataCharacteristicUUID, err = gatt.ParseUUID(DataCharacteristicString)
	if err != nil {
		log.Fatalf("error parsing DataCharacteristicString: %v", err)
		return
	}

	HCIDevice, err = gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}
}

func LogPeripheralData(p gatt.Peripheral) {
	str := "Peripheral\n\tID       = %s\n\tPeripheral Name     = %s\n\tPeripheral Services = %s\n"
	global.Debugf(str, p.ID(), p.Name(), p.Services())
}

func LogAdvertisementData(a *gatt.Advertisement) {
	str := "Advertisement\n\tLocalName        = %s\nAdvertisement ManufacturerData = %s\n\tServiceData      = %s\n\tServices         = %s\n\tOverflowService  = %s\n\tTxPowerLevel     = %s\n\tConnectable      = %s\n\tSolicitedService = %s"
	global.Debugf(str, a.LocalName, a.ManufacturerData, a.ServiceData, a.Services, a.OverflowService, a.TxPowerLevel, a.Connectable, a.SolicitedService)
}
