package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Id            string
	ProductCode   string
	AdminName     string
	OrderQuantity int
	TotalPrice    int
}

type OrderWithProduct struct {
	Order
	Product
}

func (a Order) Validate() error {
	return validation.ValidateStruct(&a,
		// ProductCode cannot be empty, and the length must between 3 and 10
		validation.Field(&a.ProductCode, validation.Required, validation.Length(3, 10)),
		// AdminName cannot be empty, and the length must between 5 and 50
		validation.Field(&a.AdminName, validation.Required, validation.Length(5, 50)),
		// OrderQuantity cannot be empty
		validation.Field(&a.OrderQuantity, validation.Required),
	)
}
