package typ

import "time"

type Account struct {
	UUID      string
	Balance   int
	Closed    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
