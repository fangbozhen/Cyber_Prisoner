package main

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Msg: "pong",
	}
	data := convertToJson(message)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
