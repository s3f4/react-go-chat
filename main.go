package main

import (
	"net/http"
)

// Channel comment
type Channel struct {
	Id   string
	Name string
}

func main() {
	router := NewRouter()
	router.Handle("channel add", addChannel)
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
