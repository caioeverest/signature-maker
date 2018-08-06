package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	
	"github.com/caioever/signature-maker/handler"
	"github.com/caioever/signature-maker/config"
	)

func API(cfg *config.Config) *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signature", handler.CreateNewSignature).Methods("POST")

	n := negroni.New()
	n.UseHandler(router)

	return n
}
