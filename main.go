package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rjrobinson/go-programming-blueprints/trace"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of this application.")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.NewTracer(os.Stdout)
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)

	// get the room going
	go r.run()
	// start the webserver
	log.Println("Starting on addr", *addr)
	http.ListenAndServe(*addr, nil)
}
