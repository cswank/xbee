package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cswank/xbee"
	serial "go.bug.st/serial.v1"
)

var (
	verbose bool
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("you must pass in the /dev/tty port")
	}
	verbose = len(args) > 1 && args[1] == "-v"

	mode := &serial.Mode{}
	port, err := serial.Open(args[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := xbee.ReadMessage(port)
		a, err := msg.GetAnalog()
		if err != nil {
			log.Fatal(err)
		}

		d, err := msg.GetDigital()
		if err != nil {
			log.Fatal(err)
		}

		addr := msg.GetAddr()
		for k, v := range a {
			fmt.Printf("%s - %s: %.2f\n", addr, k, v)
		}
		for k, v := range d {
			fmt.Printf("%s - %s: %t\n", addr, k, v)
		}
	}
}
