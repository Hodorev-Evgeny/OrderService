package feature_order_service

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
	pr "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

type OrderRepository interface {
	CreateOrder(
		ctx context.Context,
		order core_domain.Order,
	) (core_domain.Order, error)

	GetOrder(
		ctx context.Context,
		OrderID int64,
	) (core_domain.Order, error)

	GetAllOrders(
		ctx context.Context,
	) ([]core_domain.Order, error)

	CancelOrder(
		ctx context.Context,
		OrderID int64,
	) (core_domain.Order, error)
}

type OrderService struct {
	repository OrderRepository
	client     pr.ProductServiceClient
}

func NewOrderService(
	repository OrderRepository,
	client pr.ProductServiceClient,
) *OrderService {
	return &OrderService{
		repository: repository,
		client:     client,
	}
}
