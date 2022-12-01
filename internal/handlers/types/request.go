package types

import (
	"customer/internal"
)

type CreateRequest struct {
	Name        string           `json:"name"`
	Address     customer.Address `json:"address"`
	PhoneNumber string           `json:"phone_number"`
	Mail        string           `json:"mail"`
	Age         int              `json:"age"`
}
