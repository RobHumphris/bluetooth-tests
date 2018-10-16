package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RobHumphris/bluetooth-tests/data"
)

func Peripherals(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(data.DiscoveredPeripherals)
	writeBytes(w, b)
}
