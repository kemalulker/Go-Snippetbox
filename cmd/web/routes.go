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

	mux.Get("/user/signup", dynamicMiddleWare.ThenFunc(app.signUpUserForm))
	mux.Post("/user/signup", dynamicMiddleWare.ThenFunc(app.signUpUser))
	mux.Get("/user/login", dynamicMiddleWare.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleWare.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleWare.ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
