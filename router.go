package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules   map[string]Handler
	session *r.Session
}

// NewRouter return new router address
func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

// FindHandler
func (r *Router) FindHandler(messageName string) (Handler, bool) {
	handler, found := r.rules[messageName]
	fmt.Printf("MessageName: %+v\n", messageName)
	fmt.Printf("%#v\n", handler)
	fmt.Println(found)
	return handler, found
}

// Handle
func (r *Router) Handle(messageName string, handler Handler) {
	fmt.Println(reflect.TypeOf(messageName))
	r.rules[messageName] = handler
	fmt.Printf("rules : %#v\n", r.rules)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("bburada bir sorun var")
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err.Error())
		return
	}

	client := NewClient(socket, r.FindHandler, r.session)
	defer client.socket.Close()
	go client.Write()
	client.Read()
}
