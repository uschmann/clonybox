package rfidreader

type RfidReader interface {
	StartReading(ch chan string)
}
