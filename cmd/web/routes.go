package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleWare := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleWare.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleWare.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleWare.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleWare.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
