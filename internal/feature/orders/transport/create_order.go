package feature_order_transport

import (
	"context"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/order"
)

func (t *OrderTransport) CreateOrder(
	ctx context.Context,
	req *pb.OrderRequest,
) (*pb.OrderResponse, error) {
	order_domain, err := t.service.CreateOrder(ctx, req.Quantity, req.ProductId.Id)
	if err != nil {
		return nil, err
	}

	response := core_domain.DomainFromResponse(order_domain)
	return response, nil
}
