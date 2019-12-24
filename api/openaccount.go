package api

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/lmuench/gommanded/api/account/event"
	"github.com/lmuench/gommanded/typ"
)

func OpenAccount(account typ.Account) (*typ.Account, error) {
	if account.Balance < 0 {
		return nil, errors.New("Initial account balance cannot be negative")
	}
	accountOpened := event.AccountOpened{
		AccountUUID:    uuid.Must(uuid.NewV4()).String(),
		InitialBalance: account.Balance,
		TimeSent:       time.Now(),
	}
	publish(accountOpened)
	return &typ.Account{
		UUID:    accountOpened.AccountUUID,
		Balance: accountOpened.InitialBalance,
		Closed:  false,
	}, nil
}
