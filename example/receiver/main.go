package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wrfly/simorse/morse"
)

func receive(index int) {
	cmdChan := morse.ReceiveSignal(context.Background())
	for x := range cmdChan {
		if x != "" {
			log.Printf("#%d received: %s", index, x)
		}
		if x == "q" {
			break
		}
	}
}

func main() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)

	go receive(1)
	go receive(2)
	go receive(3)
	receive(4)
}
