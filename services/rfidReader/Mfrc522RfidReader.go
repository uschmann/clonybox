package rfidreader

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/mfrc522"
	"periph.io/x/host/v3"
)

type Mfrc522RfidReader struct {
	rfid     *mfrc522.Dev
	port     spi.PortCloser
	lastRfid string
	resetPin string
	irqPin   string
}

func NewMfrc522RfidReader(resetPin string, irqPin string) *Mfrc522RfidReader {
	mfrc522RfidReader := &Mfrc522RfidReader{
		resetPin: resetPin,
		irqPin:   irqPin,
	}
	var err error

	// guarantees all drivers are loaded.
	if _, err = host.Init(); err != nil {
		log.Fatal(err)
	}

	// get the first available spi port eith empty string.
	mfrc522RfidReader.port, err = spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	// get GPIO rest pin from its name
	var gpioResetPin gpio.PinOut = gpioreg.ByName(mfrc522RfidReader.resetPin)
	if gpioResetPin == nil {
		log.Fatalf("Failed to find %v", mfrc522RfidReader.resetPin)
	}

	// get GPIO irq pin from its name
	var gpioIRQPin gpio.PinIn = gpioreg.ByName(mfrc522RfidReader.irqPin)
	if gpioIRQPin == nil {
		log.Fatalf("Failed to find %v", mfrc522RfidReader.irqPin)
	}

	mfrc522RfidReader.rfid, err = mfrc522.NewSPI(mfrc522RfidReader.port, gpioResetPin, gpioIRQPin, mfrc522.WithSync())
	if err != nil {
		log.Fatal(err)
	}

	// setting the antenna signal strength, signal strength from 0 to 7
	mfrc522RfidReader.rfid.SetAntennaGain(7)

	fmt.Println("Started rfid reader.")

	return mfrc522RfidReader
}

func (m *Mfrc522RfidReader) Close() {
	if err := m.rfid.Halt(); err != nil {
		log.Fatal(err)
	}

	if err := m.port.Close(); err != nil {
		log.Fatal(err)
	}
}

func (m *Mfrc522RfidReader) StartReading(ch chan RfidEvent) {
	for {
		data, err := m.rfid.ReadUID(1 * time.Second)
		if err == nil {
			rfid := hex.EncodeToString(data)

			if rfid != m.lastRfid {
				m.lastRfid = rfid
				ch <- RfidEvent{
					Type: "scanned",
					Rfid: rfid,
				}
			}
		}
	}
}
