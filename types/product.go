package types

type Product struct {
	ID          string `json:"_id"`
	SKU         string `json:"sku"`
	ProductName string `json:"product_name"`
}
