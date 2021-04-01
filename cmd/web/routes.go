package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippetForm).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")
	r.HandleFunc("/snippet/{id}", app.showSnippet).Methods("GET")

	n := negroni.New()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	r.PathPrefix("/static/").Handler(n.With(
		negroni.Wrap(http.StripPrefix("/static/", fileServer)),
	)).Methods("GET", "OPTIONS")

	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	return n
}
