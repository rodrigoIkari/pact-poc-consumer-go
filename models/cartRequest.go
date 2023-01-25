package models

type CartRequest struct {
	Items            []CartItem `json:"items" validate:""`
	Buyer_Asset_Code string     `json:"buyer-asset-code"  validate:"required"`
}

type CartItem struct {
	SKU               string  `json:"sku" validade:"required"`
	Quantity          int     `json:"quantity" validade:"required"`
	Amount            float64 `json:"amount" validade:"required"`
	Seller_Asset_Code string  `json:"seller-asset-code" validade:"required"`
}
