package main

import (
	"flag"
	"groupie-tracker/web"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	web.Web(addr)
}
