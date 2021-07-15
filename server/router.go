package server

import (
	"github.com/gorilla/mux"
	"github.com/mylex/go-lang-hb-test/app/handler"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.RequestHandler)

	return router
}
