package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	morse "github.com/wrfly/simorse"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you must provide a valid PID")
	}

	target := os.Args[1]
	targetPID, err := strconv.Atoi(target)
	if err != nil {
		log.Fatalf("convert pid error: %s", err)
	}

	cmd := ""
	for {
		fmt.Print("->")
		fmt.Scanln(&cmd)
		log.Printf("[%s]", cmd)
		if cmd != "" {
			morse.SendSignal(targetPID, cmd)
		}
		if cmd == "q" {
			break
		}
	}
}
