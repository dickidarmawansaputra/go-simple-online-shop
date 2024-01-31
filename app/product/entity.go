package product

import (
	"time"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/google/uuid"
)

type Product struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func NewProductPaginationFromProductRequest(req ListProductRequestPayload) ProductPagination {
	req = req.GenerateDefaultValue()
	return ProductPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func NewProductFromCreateProductRequest(req ProductRequestPayload) Product {
	return Product{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Stock:     req.Stock,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p Product) Validate() (err error) {
	if err = p.ValidateName(); err != nil {
		return
	}

	if err = p.ValidateStock(); err != nil {
		return
	}

	if err = p.ValidatePrice(); err != nil {
		return
	}

	return
}

func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductRequired
	}
	if len(p.Name) < 3 {
		return response.ErrProductInvalid
	}

	return
}

func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}

	return
}

func (p Product) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}

	return
}

func (p Product) ToProductListResponse() ProductListResponse {
	return ProductListResponse{
		Id:    p.Id,
		SKU:   p.SKU,
		Name:  p.Name,
		Stock: p.Stock,
		Price: p.Price,
	}
}

func (p Product) ToProductDetailResponse() ProductDetailResponse {
	return ProductDetailResponse{
		Id:        p.Id,
		SKU:       p.SKU,
		Name:      p.Name,
		Stock:     p.Stock,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
