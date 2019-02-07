package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// FindHandler func
type FindHandler func(string) (Handler, bool)

// Client is contain message object.
type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
}

// Message is message
type Message struct {
	Name string
	Data interface{}
}

// NewClient return new client pointer
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		fmt.Printf("%#v", msg)
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}
