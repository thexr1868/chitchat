package main

import (
	"github.com/thexr1868/chitchat/data"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Thread()

}
