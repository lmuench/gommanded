package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/gommanded/api"
	"github.com/lmuench/gommanded/typ"
	"github.com/lmuench/gommanded/web/view"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var accountParams typ.Account
	err := json.NewDecoder(r.Body).Decode(&accountParams)
	if err != nil {
		respondWithError(w, err)
		return
	}
	createdAccount, err := api.OpenAccount(accountParams)
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithAccount(w, createdAccount)
}

func respondWithAccount(w http.ResponseWriter, account *typ.Account) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(view.AccountJSON(account))
}

func respondWithError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(view.ErrorJSON(err))
}
