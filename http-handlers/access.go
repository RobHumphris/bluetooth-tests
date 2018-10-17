package httphandlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RobHumphris/bluetooth-tests/bluetooth-functions"
	"github.com/RobHumphris/bluetooth-tests/global"
	"github.com/gorilla/mux"
)

var accessChannel = make(chan string)

const resultJSON = "{\"result\": \"%s\"}"
const errorJSON = "{\"error\": \"%s\"}"

func Access(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mac := vars["mac"]
	global.Debugf("Accessing %s\n", mac)
	bluetoothfunctions.Accessor(mac, accessChannel)
	global.Debugf("Waiting for access response")
	select {
	case result := <-accessChannel:
		global.Debugf("Access response")
		writeBytes(w, []byte(fmt.Sprintf(resultJSON, result)))
	case <-time.After(5 * time.Second):
		global.Debugf("Error response")
		writeErrorResponse(w, fmt.Errorf("Timeout Accessing device: %s", mac))
	}
}
