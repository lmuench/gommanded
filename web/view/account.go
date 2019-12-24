package view

import (
	"encoding/json"
	"log"

	"github.com/lmuench/gommanded/typ"
)

func status(closed bool) string {
	if closed {
		return "closed"
	}
	return "open"
}

type accountView struct {
	UUID    string
	Balance int
	Status  string
}

func AccountJSON(account *typ.Account) []byte {
	b, err := json.Marshal(accountView{
		UUID:    account.UUID,
		Balance: account.Balance,
		Status:  status(account.Closed),
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}
