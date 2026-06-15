package feature_order_transport

import (
	"context"
	"fmt"

	core_domain "github.com/Hodorev-Evgeny/OrderService/internal/core/domain"
	pb "github.com/Hodorev-Evgeny/inventory-system-api/api/order"
)

func (t *OrderTransport) GetAllOrders(
	ctx context.Context,
) (*pb.ListOrder, error) {

	list, err := t.service.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get all orders: %w", err)
	}

	response := ListDomainToResponse(list)
	return response, nil
}

func ListDomainToResponse(list []core_domain.Order) *pb.ListOrder {
	var newList pb.ListOrder
	for _, order := range list {
		newList.ListOrder = append(newList.ListOrder, core_domain.DomainFromResponse(order))
	}

	return &newList
}
