package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of this application.")
	flag.Parse()

	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()
	// start the webserver
	log.Println("Starting on addr", *addr)
	http.ListenAndServe(*addr, nil)
}
