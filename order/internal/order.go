package order

import "time"

type Order struct {
	Id         string
	CustomerId string
	Quantity   int
	TotalPrice float64
	Address    Address
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Address struct {
	City     string
	CityCode int
	PostCode int
}
