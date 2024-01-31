package transaction

import "github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"

type Product struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock uint8  `json:"stock"`
	Price int    `json:"price"`
}

func (p Product) IsExists() bool {
	return p.Id != 0
}

func (p *Product) UpdateStockProduct(amount uint8) (err error) {
	if p.Stock < uint8(amount) {
		return response.ErrAmountGreaterThanStock
	}

	p.Stock = p.Stock - uint8(amount)

	return
}
