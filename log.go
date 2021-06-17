package simorse

import (
	"log"
	"os"
)

var (
	_debug = os.Getenv("SIMORSE_DEBUG") == "true"
	l      = log.New(os.Stdout, "simorse ", 0)
)

func debug(format string, v ...interface{}) {
	if _debug {
		l.Printf(format, v...)
	}
}
