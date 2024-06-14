package web

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/artist", app.Artist)

	fileServer := http.FileServer(http.Dir("./ui/css/"))
	mux.Handle("/css/", http.StripPrefix("/css", fileServer))

	return mux
}
