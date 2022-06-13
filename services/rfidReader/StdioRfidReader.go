package rfidreader

import (
	"bufio"
	"os"
	"strings"
)

type StdioRfidReader struct {
	reader *bufio.Reader
}

func NewStdioRfidReader() *StdioRfidReader {
	return &StdioRfidReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *StdioRfidReader) StartReading(ch chan RfidEvent) {
	for {
		text, _ := r.reader.ReadString('\n')
		rfid := strings.ReplaceAll(text, "\n", "")
		ch <- RfidEvent{Type: EVENT_SCANNED, Rfid: rfid}
	}
}
