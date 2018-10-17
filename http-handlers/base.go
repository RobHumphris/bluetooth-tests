package httphandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RobHumphris/bluetooth-tests/data"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
)

var Mux *mux.Router
var Stats *stats.Stats
var StatsMiddleware http.Handler

func init() {
	Mux = mux.NewRouter()
	Stats = stats.New()
	StatsMiddleware = Stats.Handler(Mux)
}

func SetupHandlers() {
	Mux.HandleFunc("/discovered", func(w http.ResponseWriter, r *http.Request) {
		writeObjectResponse(data.DiscoveredPeripherals.List(), w)
	}).Methods("GET")

	Mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeObjectResponse(Stats.Data(), w)
	}).Methods("GET")

	Mux.HandleFunc("/access/{mac}", Access).Methods("GET")

	Mux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./bin/static/"))))
}

func writeHeader(w http.ResponseWriter, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if status != nil {
		w.WriteHeader(status[0])
	}
}

func writeBytes(w http.ResponseWriter, b []byte) {
	writeHeader(w)
	w.Write(b)
}

func writeErrorResponse(w http.ResponseWriter, e error) {
	str := fmt.Sprintf("{ \"error\": \"%s\" }", e.Error())
	writeBytes(w, []byte(str))
}

func writeObjectResponse(o interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(o)

	if err != nil {
		writeHeader(w, http.StatusInternalServerError)
	} else {
		writeBytes(w, json)
	}
}
