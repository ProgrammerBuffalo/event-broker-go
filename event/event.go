package event

import "time"

type Event struct {
	Payload   string
	Timestamp time.Time
}

func NewEvent(payload string, timestamp time.Time) *Event {
	return &Event{
		Payload:   payload,
		Timestamp: timestamp,
	}
}
