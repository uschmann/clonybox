package services

import (
	"bufio"
	"os"
	"strings"
)

type RfidReader struct {
	reader *bufio.Reader
}

func NewRfidReader() *RfidReader {
	return &RfidReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *RfidReader) StartReading(ch chan string) {
	for {
		text, _ := r.reader.ReadString('\n')
		ch <- strings.ReplaceAll(text, "\n", "")
	}
}
