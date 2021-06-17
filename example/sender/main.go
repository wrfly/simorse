package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wrfly/simorse/morse"
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

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("->")
	for scanner.Scan() {
		cmd := scanner.Text()
		sent, err := morse.SendSignal(targetPID, cmd)
		if err != nil {
			log.Printf("send signal error: %s", err)
			continue
		}
		log.Printf("[%s]", sent)
		if cmd == "q" {
			break
		}
		fmt.Print("->")
	}
}
