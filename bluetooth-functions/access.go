package bluetoothfunctions

import (
	"fmt"

	"github.com/RobHumphris/bluetooth-tests/data"
	"github.com/RobHumphris/bluetooth-tests/global"
	"github.com/paypal/gatt"
)

var accessChannel chan string

func Accessor(mac string, ac chan string) {
	accessChannel = ac
	HCIDevice.StopScanning()

	HCIDevice.Handle(
		gatt.PeripheralConnected(onPeriphConnected),
		gatt.PeripheralDisconnected(onPeriphDisconnected),
	)

	var peripheral = data.DiscoveredPeripherals.Get(mac)
	HCIDevice.Connect(peripheral)
}

func cleanUp(p gatt.Peripheral) {
	global.Debugf("cleanUp invoked")
	HCIDevice.CancelConnection(p)
	accessChannel <- p.ID()
}

func onPeriphConnected(p gatt.Peripheral, err error) {
	global.Debugf("%s connected\n", p.ID())
	defer cleanUp(p)

	services, err := p.DiscoverServices(nil)
	if err != nil {
		global.Debugf("Failed to discover service %s, err: %s\n", GenericAccessServiceUUID, err)
		return
	}

	data, err := extractCharacteristicData(GenericAccessServiceUUID, DeviceNameUUID, p, services)
	if err != nil {
		global.Debugf("Failed to extract data for service UUID %s, characteristics %s, err: %s\n", GenericAccessServiceUUID.String(), DeviceNameUUID.String(), err)
		return
	}

	deviceName := string(data[:])
	global.Debugf("DeviceName: %s\n", deviceName)

	// Now lets try writing to this bad boy:
	characteristic, err := extractCharacteristic(AvertisedServiceUUID, CommandCharacteristicUUID, p, services)
	if err != nil {
		global.Debugf("Failed to extract data for service UUID %s, characteristics %s, err: %s\n", GenericAccessServiceUUID.String(), DeviceNameUUID.String(), err)
		return
	}

	// will need to write to Descriptor: 2902        Name: Client Characteristic Configuration
	// Descriptor read value: "\x00\x00"
	// As this should set the CCCD

	/*
		BluetoothGattDescriptor descriptor = characteristic.getDescriptor(UUID.fromString(SampleGattAttributes.CLIENT_CHARACTERISTIC_CONFIG));
		descriptor.setValue(BluetoothGattDescriptor.ENABLE_NOTIFICATION_VALUE);
		mBluetoothGatt.writeDescriptor(descriptor);
	*/

	//WriteDescriptor(d *Descriptor, b []byte) error

	if (characteristic.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
		err := p.SetNotifyValue(characteristic, func(c *gatt.Characteristic, b []byte, err error) {
			global.Debugf("Notified: % X | %q\n", b, b)
		})

		if err != nil {
			global.Debugf("Failed to subscribe characteristic, err: %s\n", err)
			return
		}
	}

	VEH_CMD_UNLOCK := []byte{00, 41, 42, 43}
	err = p.WriteCharacteristic(characteristic, VEH_CMD_UNLOCK, false)
	if err != nil {
		global.Debugf("Failed to write unlock to command characteristic")
	}

	//GetAllPeripheralData(p)
}

func getService(serviceUUID gatt.UUID, services []*gatt.Service) (*gatt.Service, error) {
	for _, s := range services {
		if s.UUID().Equal(serviceUUID) {
			return s, nil
		}
	}
	return nil, fmt.Errorf("Could not find Service UUID %s in Discovered Services", serviceUUID.String())
}

func getCharacteristic(characteristicUUID gatt.UUID, characteristics []*gatt.Characteristic) (*gatt.Characteristic, error) {
	for _, c := range characteristics {
		if c.UUID().Equal(characteristicUUID) {
			return c, nil
		}
	}
	return nil, fmt.Errorf("Could not find Characteristic UUID %s in Discovered Characteristics", characteristicUUID.String())
}

func extractCharacteristic(serivceUUID gatt.UUID, characteristicUUID gatt.UUID, p gatt.Peripheral, services []*gatt.Service) (*gatt.Characteristic, error) {
	service, err := getService(serivceUUID, services)
	if err != nil {
		return nil, fmt.Errorf("Failed to find service UUID %s, err: %s", serivceUUID, err)
	}

	characteristics, err := p.DiscoverCharacteristics(nil, service)
	if err != nil {
		return nil, fmt.Errorf("Failed to obtain Service's characteristics, err: %s", err)
	}

	characteristic, err := getCharacteristic(characteristicUUID, characteristics)
	if err != nil {
		return nil, fmt.Errorf("Failed to obtain characteristic with UUID %s, err %s", characteristicUUID.String(), err)
	}
	return characteristic, nil
}

func extractCharacteristicData(serivceUUID gatt.UUID, characteristicUUID gatt.UUID, p gatt.Peripheral, services []*gatt.Service) ([]byte, error) {
	characteristic, err := extractCharacteristic(serivceUUID, characteristicUUID, p, services)
	if err != nil {
		return nil, err
	}

	byteArray, err := p.ReadCharacteristic(characteristic)
	if err != nil {
		return nil, fmt.Errorf("Could not read characteristic, err %s", err)
	}
	return byteArray, nil
}

func onPeriphDisconnected(p gatt.Peripheral, err error) {
	global.Debugf("%s disconnected\n", p.ID())
}

/**
Effectively Test data - good for finding stuff
*/
func GetAllPeripheralData(p gatt.Peripheral) {
	ss, err := p.DiscoverServices(nil)
	if err != nil {
		global.Debugf("Failed to discover services, err: %s\n", err)
		return
	}

	for _, s := range ss {
		global.Debugf("Service: %s\tName: %s\n", s.UUID().String(), s.Name())

		cs, err := p.DiscoverCharacteristics(nil, s)
		if err != nil {
			global.Debugf("Failed to discover characteristics, err: %s\n", err)
			continue
		}

		for _, c := range cs {
			global.Debugf("Characteristic: %s\tName: %s\t Properties: %s\n", c.UUID().String(), c.Name(), c.Properties().String())
			if (c.Properties() & gatt.CharRead) != 0 {
				b, err := p.ReadCharacteristic(c)
				if err != nil {
					global.Debugf("Failed to read characteristic, err: %s\n", err)
					continue
				}
				global.Debugf("Characteristic read value: %q\n\n", b)
			}

			ds, err := p.DiscoverDescriptors(nil, c)
			if err != nil {
				global.Debugf("Failed to discover descriptors, err: %s\n", err)
				continue
			}

			for _, d := range ds {
				global.Debugf("Descriptor: %s\tName: %s\n", d.UUID().String(), d.Name())
				b, err := p.ReadDescriptor(d)
				if err != nil {
					global.Debugf("Failed to read descriptor, err: %s\n", err)
					continue
				}
				global.Debugf("Descriptor read value: %q\n\n", b)
			}

			// Subscribe the characteristic, if possible.
			if (c.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
				f := func(c *gatt.Characteristic, b []byte, err error) {
					global.Debugf("Notified: % X | %q\n", b, b)
				}
				if err := p.SetNotifyValue(c, f); err != nil {
					global.Debugf("Failed to subscribe characteristic, err: %s\n", err)
					continue
				}
			}

		}
		global.Debugf("\n")
	}
}
