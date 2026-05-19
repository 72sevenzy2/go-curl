package main

import (
	"net/http"
	"testing"
)

func boolre() *bool {
	b := true
	return &b
}

func TestLogger(t *testing.T) {
	cl, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}

	Log(&client, cl, boolre(), 0)
}