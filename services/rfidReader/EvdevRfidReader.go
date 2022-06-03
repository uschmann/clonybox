package rfidreader

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

type EvdevRfIdReader struct {
	Filename *string
	device   *evdev.InputDevice
}

var scancodes = map[evdev.EvCode]string{
	evdev.KEY_0: "0",
	evdev.KEY_1: "1",
	evdev.KEY_2: "2",
	evdev.KEY_3: "3",
	evdev.KEY_4: "4",
	evdev.KEY_5: "5",
	evdev.KEY_6: "6",
	evdev.KEY_7: "7",
	evdev.KEY_8: "8",
	evdev.KEY_9: "9",
}

func NewEvdevRfIdReader(filename string) *EvdevRfIdReader {
	d, err := evdev.Open(filename)
	if err != nil {
		fmt.Printf("Cannot read %s: %v\n", filename, err)
		return nil
	}

	inputID, err := d.InputID()
	if err == nil {
		fmt.Printf("Input device ID: bus 0x%x vendor 0x%x product 0x%x version 0x%x\n",
			inputID.BusType, inputID.Vendor, inputID.Product, inputID.Version)
	}

	name, err := d.Name()
	if err == nil {
		fmt.Printf("Input device name: \"%s\"\n", name)
	}

	phys, err := d.PhysicalLocation()
	if err == nil {
		fmt.Printf("Input device physical location: %s\n", phys)
	}

	return &EvdevRfIdReader{
		Filename: &filename,
		device:   d,
	}
}

func (r *EvdevRfIdReader) StartReading(ch chan string) {
	for {
		ch <- r.readString()
	}
}

func (r *EvdevRfIdReader) readString() string {
	str := ""
	for {
		e, err := r.device.ReadOne()
		if err != nil {
			fmt.Printf("Error reading from device: %v\n", err)
			return ""
		}

		if e.Type == evdev.EV_KEY && e.Value == 1 {
			if e.Code == evdev.KEY_ENTER {
				return str
			}

			character := scancodes[e.Code]
			str += character
		}
	}
}
