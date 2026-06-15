package feature_order_service

import (
	"context"
	"errors"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (s *OrderService) GetOrder(
	ctx context.Context,
	OrderID int64,
) (core_domain.Order, error) {
	if OrderID == core_domain.UnanalyzedID {
		return core_domain.Order{}, errors.New("Order ID is un-analyzed")
	}

	order, err := s.repository.GetOrder(ctx, OrderID)
	if err != nil {
		return core_domain.Order{}, err
	}
	return order, nil
}
