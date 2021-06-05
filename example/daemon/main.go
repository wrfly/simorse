package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wrfly/simorse"
)

func hello() error {
	fmt.Println("hello")
	return nil
}

func main() {
	log.Printf("I'm a daemon program with pid=%d", os.Getpid())

	simorse.Register("h", hello)

	// block here
	var c chan bool
	<-c
}
