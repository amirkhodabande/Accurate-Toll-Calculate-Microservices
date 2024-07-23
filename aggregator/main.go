package main

import (
	"aggregator/handlers"
	"aggregator/store"
	"fmt"
	"net/http"
)

func main() {
	storage := store.NewMemoryStore()
	aggregateHandler := handlers.NewAggregateHandler(storage)
	invoiceHandler := handlers.NewInvoiceHandler(storage)

	fmt.Println("http transport is running...")
	http.HandleFunc("/aggregate", aggregateHandler.Handle)
	http.HandleFunc("/invoice", invoiceHandler.Handle)
	http.ListenAndServe(":3000", nil)

	fmt.Println("aggregator working")
}
