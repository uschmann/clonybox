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

func (r *StdioRfidReader) StartReading(ch chan string) {
	for {
		text, _ := r.reader.ReadString('\n')
		ch <- strings.ReplaceAll(text, "\n", "")
	}
}
