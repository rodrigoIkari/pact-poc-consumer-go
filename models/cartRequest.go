package models

type CartRequest struct {
	Items            []CartItem `json:"items"`
	Buyer_Asset_Code string     `json:"buyer-asset-code"`
}

type CartItem struct {
	SKU               string  `json:"sku"`
	Quantity          int     `json:"quantity"`
	Amount            float64 `json:"amount"`
	Seller_Asset_Code string  `json:"seller-asset-code"`
}
