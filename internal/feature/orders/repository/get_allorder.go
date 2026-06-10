package feature_order_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (r *OrderRepository) GetAllOrders(
	ctx context.Context,
) ([]core_domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		SELECT id, product_id, quantity, total_price, status, time_created
		FROM service_order.orders
		`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error from repository: %w", err)
	}

	defer rows.Close()
	var orders []core_domain.Order
	for rows.Next() {
		var domain_order core_domain.Order
		err := rows.Scan(
			&domain_order.ID,
			&domain_order.ProductID,
			&domain_order.Quantity,
			&domain_order.Total,
			&domain_order.Status,
			&domain_order.TimeCreated,
		)
		if err != nil {
			return nil, fmt.Errorf("error from repository: %w", err)
		}
		orders = append(orders, domain_order)
	}

	return orders, nil
}
