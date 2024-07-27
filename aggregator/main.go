package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"aggregator/handlers"
	pb "aggregator/proto"
	pbInvoicer "aggregator/proto/invoicer"
	"aggregator/services/grpc_server"
	"aggregator/store"

	"google.golang.org/grpc"
)

func main() {
	storage := store.NewMemoryStore()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		aggregatorGrpcServer := grpc_server.NewAggregatorGRPCServer(storage)
		pb.RegisterAggregatorServer(s, aggregatorGrpcServer)

		invoicerGrpcServer := grpc_server.NewInvoicerGRPCServer(storage)
		pbInvoicer.RegisterInvoicerServer(s, invoicerGrpcServer)

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	aggregateHandler := handlers.NewAggregateHandler(storage)
	invoiceHandler := handlers.NewInvoiceHandler(storage)

	fmt.Println("http transport is running...")
	http.HandleFunc("/aggregate", aggregateHandler.Handle)
	http.HandleFunc("/invoice", invoiceHandler.Handle)
	http.ListenAndServe(":3000", nil)

	fmt.Println("aggregator working")
}
