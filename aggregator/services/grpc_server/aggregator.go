package grpc_server

import (
	"context"
	"log"

	pb "aggregator/proto"
	"aggregator/services"
	"aggregator/store"
	"aggregator/types"
)

type GRPCServer struct {
	pb.UnimplementedAggregatorServer
	store store.Storer
}

func NewGRPCServer(store store.Storer) *GRPCServer {
	return &GRPCServer{
		store: store,
	}
}

func (s *GRPCServer) AggregateInvoice(ctx context.Context, in *pb.AggregateRequest) (*pb.AggregateReply, error) {
	log.Printf("Received: %v", in.GetValue())
	invoiceAggregatorService := services.NewInvoiceAggregator(s.store)

	distance := types.Distance{
		OBUID: int(in.GetOBUID()),
		Value: in.GetValue(),
		Unix:  in.GetUnix(),
	}

	if err := invoiceAggregatorService.AggregateDistance(distance); err != nil {
		return nil, err
	}

	return &pb.AggregateReply{Value: in.GetValue()}, nil
}
