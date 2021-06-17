package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/wrfly/gua"

	"github.com/wrfly/simorse/morse"
)

// TODO: simorse client
// ask target pid to connect the temp socket
// avaliable cmds:
// ls - list all the functions registered
// exec - exec one function
// need to consider stdout/err

type flags struct {
	PID int    `name:"pid" desc:"target process pid"`
	CMD string `name:"cmd" desc:"execute this cmd"`
}

var (
	socketPath = fmt.Sprintf("%s/simorse-%d.sock", os.TempDir(), os.Getpid())
)

func main() {
	flag := new(flags)
	gua.Parse(flag)

	defer os.Remove(socketPath)
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("listen socket %s err: %s", socketPath, err)
	}
	defer listener.Close()

	accept := make(chan bool)
	ready := make(chan bool)

	// send signal to target process
	go func() {
		<-accept
		cmd := fmt.Sprintf("simorse_%d", os.Getpid())
		if _, err := morse.SendSignal(flag.PID, cmd); err != nil {
			log.Fatalf("send signal to %d err: %s", flag.PID, err)
		}
		close(ready)
	}()

	// accept conn
	close(accept)
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("accept conn err: %s", err)
	}
	defer conn.Close()

	log.Printf("1")
	<-ready
	conn.Write([]byte("ping"))
	log.Printf("2")

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		log.Fatalf("read err: %s", err)
	}
	log.Printf("read: %s", bs[:n])
}
