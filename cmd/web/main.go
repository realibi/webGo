package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	//Router aka mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/messages", showMessage)
	mux.HandleFunc("/messages/create", createMessage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on port %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil{
		print(err.Error())
	}
}
