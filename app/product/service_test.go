package product

import (
	"context"
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/external/database"
	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
	"github.com/stretchr/testify/assert"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := ProductRequestPayload{
			Name:  "Product 1",
			Stock: 10,
			Price: 1000,
		}

		err := svc.CreateProduct(context.Background(), req)
		assert.Nil(t, err)
	})
	t.Run("name is required", func(t *testing.T) {
		req := ProductRequestPayload{
			Name:  "",
			Stock: 10,
			Price: 1000,
		}

		err := svc.CreateProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("stock is invalid", func(t *testing.T) {
		req := ProductRequestPayload{
			Name:  "Product 1",
			Stock: 0,
			Price: 1000,
		}

		err := svc.CreateProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrStockInvalid, err)
	})
	t.Run("price is invalid", func(t *testing.T) {
		req := ProductRequestPayload{
			Name:  "Product 1",
			Stock: 10,
			Price: 0,
		}

		err := svc.CreateProduct(context.Background(), req)
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPriceInvalid, err)
	})
}

func TestListProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pagination := ListProductRequestPayload{
			Cursor: 0,
			Size:   10,
		}

		products, err := svc.ListProduct(context.Background(), pagination)
		assert.Nil(t, err)
		assert.NotNil(t, products)
	})
}

func TestDetailProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product, err := svc.ProductDetail(context.Background(), "f0e49575-f81d-4ddd-8e7e-68f1afbc34a5")
		assert.Nil(t, err)
		assert.NotNil(t, product)
	})
}
