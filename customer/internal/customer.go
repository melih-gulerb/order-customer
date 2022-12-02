package customer

import "time"

type Customer struct {
	Id          string
	Name        string
	Address     Address
	PhoneNumber string
	Mail        string
	Age         int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Address struct {
	City     string
	CityCode int
	PostCode int
}
