package feature_order_repository

import (
	core_repository_pool "github.com/Hodorev-Evgeny/OrderService/internal/core/repository/postgres"
)

type OrderRepository struct {
	pool core_repository_pool.Pool
}

func NewOrderRepository(
	pool core_repository_pool.Pool,
) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}
