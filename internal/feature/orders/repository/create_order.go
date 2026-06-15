package feature_order_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (r *OrderRepository) CreateOrder(
	ctx context.Context,
	order core_domain.Order,
) (core_domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO service_order.orders (product_id, quantity, total_price, status, time_created)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, product_id, quantity, total_price, status, time_created;`

	row := r.pool.QueryRow(ctx, query,
		order.ProductID,
		order.Quantity,
		order.Total,
		order.Status,
		order.TimeCreated,
	)

	var domain_order core_domain.Order
	err := row.Scan(
		&domain_order.ID,
		&domain_order.ProductID,
		&domain_order.Quantity,
		&domain_order.Total,
		&domain_order.Status,
		&domain_order.TimeCreated,
	)
	if err != nil {
		return core_domain.Order{}, fmt.Errorf("error creating order: %w", err)
	}

	return domain_order, nil
}
