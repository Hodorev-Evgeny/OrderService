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
