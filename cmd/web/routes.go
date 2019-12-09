package main

import (
	"net/http"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.creteSnippet)
	mux.HandleFunc("/test", app.test)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standartMiddleware.Then(mux)
}
