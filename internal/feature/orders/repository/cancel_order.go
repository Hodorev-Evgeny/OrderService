package feature_order_repository

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
)

func (r *OrderRepository) CancelOrder(
	ctx context.Context,
	OrderID int64,
) (core_domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		UPDATE service_order.orders
		SET status = $1
		WHERE id = $2
		RETURNING id, product_id, quantity, total_price, status, time_created;`

	row := r.pool.QueryRow(ctx, query,
		core_domain.CANCELLED,
		OrderID,
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
		return core_domain.Order{}, fmt.Errorf("error cancel order: %w", err)
	}

	return domain_order, nil
}
