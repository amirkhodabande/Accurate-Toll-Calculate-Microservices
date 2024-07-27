package main

import (
	"aggregator/services/client"
	"fmt"
	"gateway/handlers"
	"net/http"
)

func main() {
	httpInvoiceGetterClient := client.NewHTTPInvoiceGetter()
	invoiceHandler := handlers.NewInvoiceHandler(httpInvoiceGetterClient)

	fmt.Println("http transport is running...")
	http.HandleFunc("/invoice", handlers.MakeApiFunc(invoiceHandler.Handle))
	http.ListenAndServe(":3001", nil)

	fmt.Println("aggregator working")
}
