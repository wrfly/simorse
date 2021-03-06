package morse

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	cmd := "abcdefghijklnmopqrstuvwxyz 1234 5sadfasds"

	cmdChan := ReceiveSignal(context.Background())
	go func() {
		x := <-cmdChan
		t.Logf("got cmd: %s", x)
		if x != cmd {
			t.Errorf("[%s] != [%s]", x, cmd)
		}
	}()

	if err := SendSignal(os.Getpid(), cmd); err != nil {
		t.Errorf("%s", err)
	}
	time.Sleep(time.Second)
}
