package main

import (
	"fmt"
	"log"
	"os"

	serial "go.bug.st/serial.v1"

	"github.com/cswank/xbee"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("you must pass in the /dev/tty port")
	}

	mode := &serial.Mode{}
	port, err := serial.Open(args[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	for {
		getDelimiter(port)
		d := make([]byte, 2)
		n, err := port.Read(d)
		if err != nil || n != 2 {
			continue
		}

		l, err := xbee.GetLength(d)
		if err != nil {
			continue

		}

		d = []byte{}
		d, err = getBody(d, port, int(l+1))
		if err != nil {
			continue
		}

		msg, err := xbee.NewMessage(d)
		if err != nil {
			continue
		}

		a, _ := msg.GetAnalog()
		for k, v := range a {
			fmt.Println(k, v)
		}
		fmt.Println("")
	}
}

func getDelimiter(port serial.Port) {
	for {
		d := make([]byte, 1)
		n, err := port.Read(d)
		if err != nil || n != 1 {
			log.Println(n, err)
			continue

		}
		if d[0] == 0x7E {
			return
		}
	}
}

func getBody(data []byte, port serial.Port, l int) ([]byte, error) {
	d := make([]byte, l)
	n, err := port.Read(d)
	if err != nil {
		return nil, err
	}
	d = append(data, d[:n]...)
	if n == l {
		return d, nil
	}
	return getBody(d, port, l-n)
}
