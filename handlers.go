package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func addChannel(client *Client, data interface{}) {
	var message Message
	var channel Channel
	mapstructure.Decode(data, &channel)
	fmt.Printf("%#v\n", channel)
	//TODO nsert into Rethinkdb
	channel.Id = "ABC123"
	message.Name = "channel add"
	message.Data = channel
	client.send <- message
}
