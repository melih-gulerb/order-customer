package storages

import (
	"github.com/google/uuid"
	order "order/internal"
	"order/internal/handlers/types"
	"time"
)

func BuildCreateOrder(request types.CreateRequest) order.Order {
	return order.Order{
		Id:         uuid.New().String(),
		CustomerId: request.CustomerId,
		Quantity:   request.Quantity,
		TotalPrice: request.TotalPrice,
		Address:    order.Address{},
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}
}
