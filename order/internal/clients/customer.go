package clients

import (
	"fmt"
	"net/http"
	"order/errors"
)

type CustomerClient struct {
	client  http.Client
	baseURL string
}

func NewCustomerClient(baseURL string) *CustomerClient {
	return &CustomerClient{
		baseURL: baseURL,
	}
}

func (c *CustomerClient) ValidateCustomerExists(customerId string) *errors.Error {
	uri := fmt.Sprintf("%s/customer/%s/validate", c.baseURL, customerId)
	_, err := c.client.Get(uri)
	if err != nil {
		return errors.Define("CustomerId is not valid: "+customerId, 400)
	}

	return nil
}
