package simorse

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/wrfly/simorse/morse"
)

func init() {
	go watchSignal()
}

// watch signals
func watchSignal() {
	cmdCh := morse.ReceiveSignal(context.Background())
	for cmd := range cmdCh {
		debug("receive cmd: %s\n", cmd)
		if !strings.HasPrefix(cmd, "simorse_") {
			fmt.Printf("unknown simorse cmd %s\n", cmd)
		}
		cliPID := strings.Split(cmd, "_")[1]
		conn, err := net.Dial("unix",
			fmt.Sprintf("%s/simorse-%s.sock", os.TempDir(), cliPID))
		if err != nil {
			log.Printf("dial simorse %s socket err: %s", cliPID, err)
			continue
		}
		// read
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			log.Printf("read socket err: %s", err)
			return
		}
		log.Printf("read %s", buff[:n])

		// write
		conn.Write([]byte("pong"))
	}
}
