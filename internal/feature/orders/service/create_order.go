package feature_order_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
	pr "github.com/Hodorev-Evgeny/inventory-system-api/api/product"
)

func (s *OrderService) CreateOrder(
	ctx context.Context,
	quantity int64,
	ProductID int64,
) (core_domain.Order, error) {
	feature, err := s.client.GetItemByID(ctx, &pr.ProductID{Id: ProductID})
	if err != nil {
		return core_domain.Order{}, fmt.Errorf("error getting product from database: %w", err)
	}
	total := feature.Price * quantity

	unanalyzed_order := core_domain.CreateOrder(
		quantity,
		ProductID,
		total,
	)

	order, err := s.repository.CreateOrder(ctx, unanalyzed_order)
	if err != nil {
		return core_domain.Order{}, fmt.Errorf("error creating order from service: %w", err)
	}

	return order, nil
}
