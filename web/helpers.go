package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type ApplicationError struct {
	Message string
	Code    int
}

var (
	templates *template.Template
	err       error
)

func init() {
	templates, err = template.ParseGlob("./ui/html/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("Ошибка при отправке ответа: %s\n", err.Error())
	app.errorLog.Output(5, trace)

	app.Errors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	app.Errors(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) BadRequest(w http.ResponseWriter) {
	app.ClientError(w, http.StatusBadRequest)
}

func (app *Application) Errors(w http.ResponseWriter, errorMessage string, errorCode int) {
	if 	err := templates.ExecuteTemplate(w, "error.html", ApplicationError {
		Message: errorMessage,
		Code:    errorCode,
	}); err != nil {
		http.Error(w, "err", http.StatusInternalServerError)
	}
}
