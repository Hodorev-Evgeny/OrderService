package feature_order_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/order"
)

func (t *OrderTransport) CancelOrder(
	ctx context.Context,
	id *pb.OrderID,
) (*pb.OrderResponse, error) {
	order, err := t.service.CancelOrder(ctx, id.Id)
	if err != nil {
		return nil, fmt.Errorf("cancel order: %w", err)
	}

	return core_domain.DomainFromResponse(order), nil
}
