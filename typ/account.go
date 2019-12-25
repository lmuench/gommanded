package typ

import "time"

type Account struct {
	id        int64 // used by Google Cloud Datastore
	CreatedAt time.Time
	UUID      string
	Balance   int
	Closed    bool
}
