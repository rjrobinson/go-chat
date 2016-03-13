package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}
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
