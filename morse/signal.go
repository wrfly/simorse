package morse

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SendSignal(pid int, cmd string) error {
	// find process
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	codes := SigMorseCode(cmd)
	for _, sig := range codes.Signal() {
		if err := process.Signal(sig.Signal()); err != nil {
			return err
		}
		// give time to let the process handle the signal
		time.Sleep(time.Millisecond * 5)
	}

	return nil
}

func ReceiveSignal(ctx context.Context) chan string {
	sigChan := make(chan os.Signal, 50)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGUSR2)

	cmdChan := make(chan string, 10)
	go func() {
		for ctx.Err() == nil {
			received := Signals{}
			tk := time.NewTicker(time.Second * 2)
			for done := false; !done; {
				select {
				case <-tk.C:
					done = true
				case s := <-sigChan:
					received = append(received, ParseSignal(s))
				}
			}
			tk.Stop()
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
