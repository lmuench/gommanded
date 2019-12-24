package view

import (
	"encoding/json"
	"log"
)

type errorView struct {
	Error string
}

func ErrorJSON(e error) []byte {
	b, err := json.Marshal(errorView{
		Error: e.Error(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}
