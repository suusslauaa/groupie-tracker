package web

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewApplication(errorLog, infoLog *log.Logger) *Application {
	return &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func NewServer(addr *string, errorLog *log.Logger, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
}

func Web(addr *string) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := NewApplication(errorLog, infoLog)
	srv := NewServer(addr, errorLog, app.Routes())

	infoLog.Printf("Запуск сервера на http://localhost%s", *addr)
	errorLog.Fatal(srv.ListenAndServe())
}
