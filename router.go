package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules map[string]Handler
}

// NewRouter return new router address
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

// FindHandler
func (r *Router) FindHandler(messageName string) (Handler, bool) {
	handler, found := r.rules[messageName]
	return handler, found
}

// Handle
func (r *Router) Handle(messageName string, handler Handler) {
	r.rules[messageName] = handler
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err.Error())
		return
	}

	client := NewClient(socket, r.FindHandler)
	go client.Write()
	client.Read()
}
