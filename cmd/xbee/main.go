package main

import (
	"fmt"
	"log"

	"github.com/cswank/xbee"
	serial "go.bug.st/serial.v1"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	port = kingpin.Arg("serial-port", "serial port from which to read").Required().String()
)

func init() {
	kingpin.Parse()
}

func main() {
	port, err := serial.Open(*port, &serial.Mode{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := xbee.ReadMessage(port)
		if err != nil {
			log.Fatal(err)
		}

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
