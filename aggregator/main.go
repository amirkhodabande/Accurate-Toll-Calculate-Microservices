package main

import (
	"aggregator/handlers"
	"aggregator/store"
	"fmt"
	"net/http"
)

func main() {
	aggregateHandler := handlers.NewAggregateHandler(store.NewMemoryStore())

	fmt.Println("http transport is running...")
	http.HandleFunc("/aggregate", aggregateHandler.Handle)
	http.ListenAndServe(":3000", nil)

	fmt.Println("aggregator working")
}
