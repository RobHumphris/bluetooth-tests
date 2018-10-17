package global

import (
	"fmt"
	"log"
	"os"
)

var debug = false

func init() {
	debug = os.Getenv("DEBUG") != "1"
	if debug {
		log.Printf("Debug Messages are On")
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		fmt.Printf(format, v...)
	}
}
