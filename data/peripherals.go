package data

import "sync"

type Peripherals struct {
	DeviceMAC string `json:"deviceMAC"`
}

type DiscoveredPeripheralsMap struct {
	sync.RWMutex
	Discovered map[string]Peripherals `json:"discovered"`
}

func NewDiscoveredPeripheralsMap() *DiscoveredPeripheralsMap {
	retObject := DiscoveredPeripheralsMap{Discovered: make(map[string]Peripherals)}
	return &retObject
}

func (dpm *DiscoveredPeripheralsMap) Store(mac string) {
	p := Peripherals{
		DeviceMAC: mac,
	}

	dpm.Lock()
	dpm.Discovered[mac] = p
	dpm.Unlock()
}
