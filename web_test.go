package main

import (
	"fmt"
	"groupie-tracker/web"
	"log"
	"os"
	"testing"
)

func TestHomeddf(t *testing.T) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := web.NewApplication(errorLog, infoLog)

	// var w http.ResponseWriter
	// var r *http.Request

	// fmt.Println(app.Home(w, r))

	fmt.Println(app.Routes())
}
