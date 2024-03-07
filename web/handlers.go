package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	if err = templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) artist(w http.ResponseWriter, r *http.Request) {
	if err = templates.ExecuteTemplate(w, "artist.html", nil); err != nil {
		app.serverError(w, err)
	}
}
