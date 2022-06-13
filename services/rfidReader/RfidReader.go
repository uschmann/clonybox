package rfidreader

type RfidReader interface {
	StartReading(ch chan RfidEvent)
}

const (
	EVENT_SCANNED = "scanned"
	EVENT_REMOVED = "removed"
)

type RfidEvent struct {
	Type string
	Rfid string
}
