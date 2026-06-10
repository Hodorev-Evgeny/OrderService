package feature_order_service

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (s *OrderService) GetAllOrders(
	ctx context.Context,
) ([]core_domain.Order, error) {
	list, err := s.repository.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("error geting list form service:", err)
	}

	return list, nil
}
