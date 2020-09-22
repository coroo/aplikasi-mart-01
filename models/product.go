package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	Id           string
	ProductCode  string
	ProductName  string
	ProductPrice int
}

type TotalProduct struct {
	Count        int64
	ProductPrice int
}

func (a Product) Validate() error {
	return validation.ValidateStruct(&a,
		// ProductCode cannot be empty, and the length must between 3 and 10
		validation.Field(&a.ProductCode, validation.Required, validation.Length(3, 10)),
		// ProductName cannot be empty, and the length must between 5 and 50
		validation.Field(&a.ProductName, validation.Required, validation.Length(5, 50)),
		// ProductPrice cannot be empty
		validation.Field(&a.ProductPrice, validation.Required),
	)
}
