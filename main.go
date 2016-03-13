package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()
	// start the webserver
	port := ":8080"
	fmt.Println("Starting on localhost", port)
	http.ListenAndServe(port, nil)
}
