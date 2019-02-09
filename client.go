package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
	"strconv"
)

// FindHandler func
type FindHandler func(string) (Handler, bool)

// Client is contain message object.
type Client struct {
	socket      *websocket.Conn
	send        chan Message
	findHandler FindHandler
	session     *r.Session
}

// Message is message
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// NewClient return new client pointer
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
		session:     session,
	}
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		fmt.Printf("%#v\n", message)
		if handler, found := client.findHandler(message.Name); found {
			fmt.Println(message.Name + strconv.FormatBool(found))
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		fmt.Print("client write")
		fmt.Printf("%#v", msg)
		if err := client.socket.WriteJSON(msg); err != nil {
			fmt.Println(err)
			break
		}
	}
	client.socket.Close()
}
