package models

import "github.com/go-playground/validator/v10"

type CartRequest struct {
	Items            []*CartItem `json:"items" validate:"required,min=1,dive"`
	Buyer_Asset_Code string      `json:"buyer-asset-code"  validate:"required,iso4217"`
}

type CartItem struct {
	SKU               string  `json:"sku" validate:"required,min=3"`
	Quantity          int     `json:"quantity" validate:"required,min=1"`
	Amount            float64 `json:"amount" validate:"required"`
	Seller_Asset_Code string  `json:"seller-asset-code" validate:"required,iso4217"`
}

func (c CartRequest) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
