package feature_order_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (service OrderService) CancelOrder(
	ctx context.Context,
	OrderID int64,
) (core_domain.Order, error) {
	if OrderID == core_domain.UnanalyzedID {
		return core_domain.Order{}, fmt.Errorf("Order ID is not analyzed")
	}

	order, err := service.repository.CancelOrder(ctx, OrderID)
	if err != nil {
		return core_domain.Order{}, fmt.Errorf("Error getting order: %w", err)
	}

	return order, nil
}
