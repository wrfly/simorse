package morse

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	cmd := "hello world 0 1 2 3 4 5"

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
