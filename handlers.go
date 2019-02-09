package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

func addChannel(client *Client, data interface{}) {
	fmt.Println("channel add Handler")
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		response, err := r.Table("channel").
			Insert(channel).
			RunWrite(client.session)
		fmt.Printf("add channel rethinkdb response : %+v\n", response)
		if err != nil {
			fmt.Println("database error")
			client.send <- Message{"error", err.Error()}
		}
	}()

}

func subscribeChannel(client *Client, data interface{}) {
	fmt.Println("---------")
	go func() {
		fmt.Printf("-------%+v\n ", "subscribeChannel")
		cursor, err := r.Table("channel").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)
		if err != nil {
			client.send <- Message{"errror", err.Error()}
		}

		var change r.ChangeResponse
		for cursor.Next(&change) {
			if change.NewValue != nil && change.OldValue == nil {
				fmt.Printf("%+v\n", change.NewValue)
				client.send <- Message{"channel add", change.NewValue}
				fmt.Println("sent channel add msg")
			}
		}
	}()
}
