package httphandlers

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	writeObjectResponse(Stats.Data(), w)
}
