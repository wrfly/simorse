package main

import (
	"context"
	"fmt"
	"log"
	"os"

	morse "github.com/wrfly/simorse"
)

func main() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)

	cmdChan := morse.ReceiveSignal(context.Background())
	for x := range cmdChan {
		if x != "" {
			log.Printf("%s", x)
		}
		if x == "q" {
			break
		}
	}
}
