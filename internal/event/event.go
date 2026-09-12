package event

import "time"

type Event struct {
	EventID string
	Type    string
	Source  string
	Payload map[string]interface{}
}

type SavedEvent struct {
	ID         int64
	ReceivedAt time.Time
}
