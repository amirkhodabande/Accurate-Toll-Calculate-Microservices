package grpc_server

import (
	"context"
	"log"

	pbInvoicer "aggregator/proto/invoicer"
	"aggregator/services"
	"aggregator/store"
)

type InvoicerGRPCServer struct {
	pbInvoicer.UnimplementedInvoicerServer
	store store.Storer
}

func NewInvoicerGRPCServer(store store.Storer) *InvoicerGRPCServer {
	return &InvoicerGRPCServer{
		store: store,
	}
}

func (s *InvoicerGRPCServer) GetInvoice(ctx context.Context, in *pbInvoicer.GetInvoiceRequest) (*pbInvoicer.GetInvoiceReply, error) {
	log.Printf("Received: %v", in.GetOBUID())
	invoiceProcessor := services.NewInvoiceProcessor(s.store)

	id := in.GetOBUID()
	invoice, err := invoiceProcessor.CalculateInvoice(int(id))
	if err != nil {
		return nil, err
	}
	return &pbInvoicer.GetInvoiceReply{
		OBUID:    int64(invoice.OBUID),
		Distance: invoice.Distance,
		Amount:   invoice.Amount,
	}, nil
}
