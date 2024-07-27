package main

import (
	"aggregator/services/client"
	"fmt"
	"gateway/handlers"
	"net/http"
)

func main() {
	// httpInvoiceGetterClient := client.NewHTTPInvoiceGetter()
	grpcInvoiceGetterClient := client.NewGRPCInvoiceGetter()
	invoiceHandler := handlers.NewInvoiceHandler(grpcInvoiceGetterClient)

	fmt.Println("http transport is running...")
	http.HandleFunc("/invoice", handlers.MakeApiFunc(invoiceHandler.Handle))
	http.ListenAndServe(":3001", nil)

	fmt.Println("aggregator working")
}
