package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	drinks map[int32]*pb.Drink
	orders map[int32]int32 // drink_id -> quantity
}

func StartGrpcServer() error {

	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{
		orders: make(map[int32]int32), // initialize empty orders map
		drinks: map[int32]*pb.Drink{ // initialize drinks map
			1: {Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
			2: {Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
			3: {Id: 3, Name: "Coffee", Price: 1, Description: "Mifare isn't that secure"},
		},
	}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions

func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinkList, error) {
	resp := &pb.DrinkList{}
	for _, d := range s.drinks {
		resp.Drinks = append(resp.Drinks, d)
	}
	return resp, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*wrapperspb.BoolValue, error) {
	for _, item := range req.Items {
		s.orders[item.DrinkId] += item.Quantity
	}
	return wrapperspb.Bool(true), nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrderTotal, error) {
	resp := &pb.OrderTotal{}
	for id, qty := range s.orders {
		resp.Totals = append(resp.Totals, &pb.OrderTotalItem{
			Drink:    s.drinks[id],
			Quantity: qty,
		})
	}
	return resp, nil
}
