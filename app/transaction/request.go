package transaction

type CreateTransactionRequestPayload struct {
	ProductSKU string `json:"product_sku"`
	Amount     uint8  `json:"amount"`
	Email      string `json:""`
}
