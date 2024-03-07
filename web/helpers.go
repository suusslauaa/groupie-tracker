package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	app.errors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	app.errors(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) errors(w http.ResponseWriter, message string, status int) {
	help := helper{
		Text: message,
		Code: status,
	}

	app.errorLog.Printf("%s: %d", message, status)

	if err = templates.ExecuteTemplate(w, "error.html", help); err != nil {
		app.serverError(w, err)
	}
}
