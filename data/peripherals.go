package data

import (
	"sync"

	"github.com/paypal/gatt"
)

type Peripherals struct {
	DeviceMAC string `json:"deviceMAC"`
}

type DiscoveredPeripheralsMap struct {
	sync.RWMutex
	Discovered map[string]gatt.Peripheral `json:"discovered"`
}

func NewDiscoveredPeripheralsMap() *DiscoveredPeripheralsMap {
	retObject := DiscoveredPeripheralsMap{Discovered: make(map[string]gatt.Peripheral)}
	return &retObject
}

func (dpm *DiscoveredPeripheralsMap) Store(p gatt.Peripheral) {
	dpm.Lock()
	dpm.Discovered[p.ID()] = p
	dpm.Unlock()
}

func (dpm *DiscoveredPeripheralsMap) List() []Peripherals {
	p := []Peripherals{}
	dpm.RLock()
	for _, value := range dpm.Discovered {
		p = append(p, Peripherals{
			DeviceMAC: value.ID(),
		})
	}
	dpm.RUnlock()
	return p
}

func (dpm *DiscoveredPeripheralsMap) Get(mac string) gatt.Peripheral {
	dpm.RLock()
	ret := dpm.Discovered[mac]
	dpm.RUnlock()
	return ret
}
