# xbee
A package for parsing automatic xbee packets.

This a single-purposed xbee package - when you have an xbee(s) that sleeps, wakes
up, and pushes messages to the coordinator the message take the form of:

    []byte{0x7E, 0x00, 0x16, 0x92, 0x00, 0x13, 0xA2, 0x00, 0x40, 0x4C, 0x0E, 0xBE, 0x61, 0x59, 0x01, 0x01, 0x00, 0x18, 0x03, 0x00, 0x10, 0x02, 0x2F, 0x01, 0xFE, 0x49}

Where

    0x7E                                      - delimiter
    0x00 0x16                                 - length (from after these two bytes until the checksum)
    0x92                                      - frame type
    0x00 0x13 0xA2 0x00 0x40 0x4C 0x0E 0xBE   - long address of the sender
    0x61 0x59                                 - short address of the sender
    0x01                                      - receive options
    0x01                                      - number of samples
    0x00 0x18                                 - digital channel mask
    0x03                                      - analog channel mask
    0x00 0x10                                 - digital samples (not present if digital channel mask == 0)
    0x02 0x2F                                 - first analog sample
    0x01 0xFE                                 - second analog sample
    0x49

To parse this message:

    func main() {
    	data := []byte{0x92, 0x00, 0x13, 0xA2, 0x00, 0x40, 0x4C, 0x0E, 0xBE, 0x61, 0x59, 0x01, 0x01, 0x00, 0x18, 0x03, 0x00, 0x10, 0x02, 0x2F, 0x01, 0xFE, 0x49}
    	x, err := xbee.NewMessage(data)
    	if err != nil {
    		log.Fatal(err)
    	}

    	a, err := x.GetAnalog()
    	if err != nil {
    		log.Fatal(err)
    	}

    	d, err := x.GetDigital()
    	if err != nil {
    		log.Fatal(err)
    	}

    	for k, v := range a {
    		fmt.Println(k, v)
    	}

    	for k, v := range d {
    		fmt.Println(k, v)
    	}
    }

The output will be:
    
    adc0: 655.7184750733138
    adc1: 598.2404692082112
    dio3: false
    dio4: true


	
