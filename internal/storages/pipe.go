package storages

import (
	c "customer/internal"
	"customer/internal/handlers/types"
	"github.com/google/uuid"
	"time"
)

func BuildCreateCustomer(request types.CreateRequest) c.Customer {
	return c.Customer{
		Id:          uuid.New().String(),
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Mail:        request.Mail,
		Age:         request.Age,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
