package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rjrobinson/go-programming-blueprints/trace"
)

type room struct {
	// forward is a ch. that holds incoming messages
	// that should be forwarded to other clients
	forward chan []byte
	// join is a chan. for users to join a chan.
	join chan *client
	// leave is a chan so users can leave
	leave chan *client
	// clients holds all clients in this room.
	clients map[*client]bool
	// tracer will recieve trace information of all activity
	tracer trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
					r.tracer.Trace("--- sent to client")
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("--- failed to send, cleaned up client")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()

	go client.write()
	client.read()
}

func newRoom() *room {
	r := &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		// nil object pattern !!!
		tracer: trace.Off(),
	}
	return r
}
