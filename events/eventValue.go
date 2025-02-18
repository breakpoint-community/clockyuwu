package events

import (
	"time"
)

type EventValue struct {
	Points         int
	Activated      bool
	ActivatedBy    string
	ActivatedAt    time.Time
	ArrivedAt      time.Time
	Partecipations map[int64]bool
}
