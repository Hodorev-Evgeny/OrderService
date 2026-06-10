package feature_order_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (s *OrderService) CreateOrder(
	ctx context.Context,
	quantity int64,
	ProductID int64,
) (core_domain.Order, error) {

	unanalyzed_order := core_domain.CreateOrder(
		quantity,
		ProductID,
		-1, // заглушка потом убрать
	)

	order, err := s.repository.CreateOrder(ctx, unanalyzed_order)
	if err != nil {
		return core_domain.Order{}, fmt.Errorf("error creating order from service: %w", err)
	}

	return order, nil
}
