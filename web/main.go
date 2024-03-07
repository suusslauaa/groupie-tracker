package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

type helper struct {
	Text string
	Code int
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

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на http://localhost%s", *addr)
	errorLog.Fatal(srv.ListenAndServe())
}
