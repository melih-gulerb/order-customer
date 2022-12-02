package types

import order "order/internal"

type CreateRequest struct {
	CustomerId string        `json:"customer_id"`
	Quantity   int           `json:"quantity"`
	TotalPrice float64       `json:"total_price"`
	Address    order.Address `json:"address"`
}
