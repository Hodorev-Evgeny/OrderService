package feature_order_service

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(
		ctx context.Context,
		order core_domain.Order,
	) (core_domain.Order, error)
}

type OrderService struct {
	repository OrderRepository
}

func NewOrderService(
	repository OrderRepository,
) *OrderService {
	return &OrderService{
		repository: repository,
	}
}
