package client

import (
	"aggregator/types"
	"context"
	"fmt"
	"log"
	"time"

	pb "aggregator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCInvoiceAggregator struct{}

func NewGRPCAggregator() *GRPCInvoiceAggregator {
	return &GRPCInvoiceAggregator{}
}

func (a *GRPCInvoiceAggregator) AggregateInvoice(distance types.Distance) (*types.AggregationClientResponse, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAggregatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AggregateInvoice(ctx, &pb.AggregateRequest{
		OBUID: int64(distance.OBUID),
		Value: distance.Value,
		Unix:  distance.Unix,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

		return nil, err
	}

	return &types.AggregationClientResponse{
		Success: true,
		Msg:     fmt.Sprintf("Aggregating: %.2f", r.GetValue()),
	}, nil
}
