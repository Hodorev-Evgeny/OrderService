package feature_order_transport

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

type OrderService interface {
	CreateOrder(
		ctx context.Context,
		quantity int64,
		ProductID int64,
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

type OrderTransport struct {
	service OrderService
}

func NewOrderTransport(
	orderRepository OrderService,
) *OrderTransport {
	return &OrderTransport{
		service: orderRepository,
	}
}
