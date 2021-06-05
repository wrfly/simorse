package morse

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SendSignal to target process
func SendSignal(pid int, cmd string) (string, error) {
	codes := SigMorseCode(cmd)
	for _, sig := range codes.Signal() {
		if err := syscall.Kill(pid, sig.Signal()); err != nil {
			return codes.String(), err
		}
		// give time to let the process handle the signal
		// TODO: try to remove this
		time.Sleep(time.Millisecond * 10)
	}

	return codes.String(), nil
}

// ReceiveSignal until ctx done
func ReceiveSignal(ctx context.Context) chan string {
	sigChan := make(chan os.Signal, 500)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGUSR2)

	cmdChan := make(chan string, 10)
	go func() {
		const tkDuration = time.Millisecond * 50
		for ctx.Err() == nil {
			received := make(Signals, 0)
			tk := time.NewTicker(tkDuration)
			for done := false; !done; {
				select {
				case <-tk.C:
					done = true
				case s := <-sigChan:
					received = append(received, ParseSignal(s))
					tk.Reset(tkDuration)
				}
			}
			tk.Stop()
			if len(received) == 0 {
				continue
			}
			sigCodes, err := ParseSigMorseCode(received.String())
			if err != nil {
				cmdChan <- fmt.Sprintf("err: %s", err)
			} else {
				cmdChan <- sigCodes.String()
			}
		}
	}()

	return cmdChan
}
