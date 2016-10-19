package xbee

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

type header struct {
	Type            uint8
	Addr            uint64
	ShortAddr       uint16
	Opts            uint8
	Samples         uint8
	DigitalChanMask uint16
	AnalogChanMask  uint8
}

type Message struct {
	header
	frame []byte
}

var (
	adc = []string{"adc0", "adc1", "adc2", "adc3"}
)

//NewMessage parses a xbee messsage (from the 3rd byte on).
//The first 3 bytes are used to detect the beginning of a new
//message and the length of the message.
func NewMessage(data []byte) (Message, error) {
	var h header
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &h)
	var msg Message
	if err != nil {
		return msg, err
	}
	msg = Message{
		header: h,
		frame:  data,
	}
	if !msg.Check() {
		err = fmt.Errorf("message failed checksum")
	}
	return msg, err
}

func (x *Message) payload() []byte {
	return x.frame[16 : len(x.frame)-1]
}

func (x *Message) GetAnalog() (map[string]float64, error) {
	m := map[string]float64{}
	if x.AnalogChanMask == 0 {
		return m, nil
	}

	var d []byte
	payload := x.payload()
	if x.DigitalChanMask == 0 {
		d = payload
	} else {
		d = payload[2:]
	}

	f := make([]uint16, len(d)/2)
	buf := bytes.NewReader(d)
	err := binary.Read(buf, binary.BigEndian, &f)
	if err != nil {
		return nil, err
	}

	var j int
	for i, o := range adc {
		if uint8(1<<uint8(i))&x.AnalogChanMask > 0 {
			x := f[j]
			m[o] = 1200.0 * float64(x) / float64(1023)
			j++
		}
	}
	return m, nil
}

func (x *Message) Check() bool {
	var total byte
	for _, item := range x.frame[:len(x.frame)-1] {
		total += item
	}
	cs := x.frame[len(x.frame)-1]
	return 0xff-total == cs
}

func (x *Message) GetAddr() string {
	return strconv.FormatUint(x.Addr, 16)
}
