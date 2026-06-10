package server

import (
	"context"
	"fmt"

	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/order"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println("Check")
	return s.orderCase.CreateOrder(ctx, req)
}

func (s *Server) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.OrderResponse, error) {
	return nil, nil
}

func (s *Server) GetAllOrder(ctx context.Context, empty *emptypb.Empty) (*pb.ListOrder, error) {
	return nil, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *pb.OrderID) (*pb.OrderResponse, error) {
	return nil, nil
}
