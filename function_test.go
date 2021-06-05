package simorse

import (
	"fmt"
	"testing"
)

func printA() error {
	fmt.Println("A")
	return nil
}

func TestSimorse(t *testing.T) {
	funcName := "print A"
	Register(funcName, printA)
	Call(funcName)
}
