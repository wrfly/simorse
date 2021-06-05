package simorse

import (
	"context"

	"github.com/wrfly/simorse/morse"
)

func init() {
	go watchSignal()
}

// watch signals
func watchSignal() {
	cmdCh := morse.ReceiveSignal(context.Background())
	for cmd := range cmdCh {
		debug("receive cmd: %s\n", cmd)
		Call(cmd)
	}
}
