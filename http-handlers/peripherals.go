package httphandlers

import (
	"net/http"

	"github.com/RobHumphris/bluetooth-tests/data"
)

func Peripherals(w http.ResponseWriter, r *http.Request) {
	//b, _ := json.Marshal(data.DiscoveredPeripherals.List())
	writeObjectResponse(data.DiscoveredPeripherals.List(), w)
	//writeBytes(w, b)
}
