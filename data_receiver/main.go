package main

import (
	"fmt"
	"net/http"

	"data-receiver.com/handlers"
)

func main() {
	wsHandler := handlers.NewWsHandler()

	http.HandleFunc("/ws", wsHandler.Handle)
	http.ListenAndServe(":30000", nil)

	fmt.Println("data_receiver working")
}
