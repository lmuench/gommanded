package event

import (
	"time"
)

type AccountOpened struct {
	AccountUUID    string
	InitialBalance int
	SentAt         time.Time
}
