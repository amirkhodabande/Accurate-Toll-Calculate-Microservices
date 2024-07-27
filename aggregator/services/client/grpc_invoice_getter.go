package client

import (
	pbInvoicer "aggregator/proto/invoicer"
	"aggregator/types"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCInvoiceGetter struct{}

func NewGRPCInvoiceGetter() *GRPCInvoiceGetter {
	return &GRPCInvoiceGetter{}
}

func (a *GRPCInvoiceGetter) GetInvoice(obuId int) (*types.InvoiceGetterClientResponse, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbInvoicer.NewInvoicerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetInvoice(ctx, &pbInvoicer.GetInvoiceRequest{
		OBUID: int64(obuId),
	})
	if err != nil {
		log.Fatalf("could not get invoice: %v", err)

		return nil, err
	}

	return &types.InvoiceGetterClientResponse{
		Success: true,
		Data: types.Invoice{
			OBUID:    int(r.GetOBUID()),
			Amount:   r.GetAmount(),
			Distance: r.GetDistance(),
		},
	}, nil
}
