package event

import "time"

type AccountOpened struct {
	AccountUUID    string
	InitialBalance int
	TimeSent       time.Time
	TimeReceived   time.Time
}
