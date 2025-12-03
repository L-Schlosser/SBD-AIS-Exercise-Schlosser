package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// todo
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResp, err := c.client.GetDrinks(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("  > id:%d name:%s price:%d description:%q\n",
			d.Id, d.Name, d.Price, d.Description)
	}
	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	_, err = c.client.OrderDrink(context.Background(), &pb.OrderRequest{
		Items: []*pb.OrderItem{
			{DrinkId: 1, Quantity: 2},
			{DrinkId: 2, Quantity: 2},
			{DrinkId: 3, Quantity: 2},
		},
	})
	if err != nil {
		return err
	}

	// 3. Order more drinks

	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	_, err = c.client.OrderDrink(context.Background(), &pb.OrderRequest{
		Items: []*pb.OrderItem{
			{DrinkId: 1, Quantity: 6},
			{DrinkId: 2, Quantity: 6},
			{DrinkId: 3, Quantity: 6},
		},
	})
	if err != nil {
		return err
	}
	// 4. Get order total

	fmt.Println("Getting the bill 💹💹💹")

	totals, err := c.client.GetOrders(context.Background(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	for _, t := range totals.Totals {
		fmt.Printf("  > Total: %d x %s\n", t.Quantity, t.Drink.Name)
	}

	fmt.Println("Orders complete!")

	return nil

	//
	// print responses after each call

}
