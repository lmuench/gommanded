package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lmuench/gommanded/web/handler"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.POST("/accounts", handler.Create)
	return router
}
