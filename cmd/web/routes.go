package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Then(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", dynamicMiddleware.Then(http.HandlerFunc(app.creteSnippet)))
	mux.Get("/snippet/:id", dynamicMiddleware.Then(http.HandlerFunc(app.showSnippet)))
	mux.Get("/test", dynamicMiddleware.Then(http.HandlerFunc(app.test)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	//return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(mux)
}
