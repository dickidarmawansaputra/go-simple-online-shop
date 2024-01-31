package product

import (
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/stretchr/testify/assert"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:  "Product 1",
			Stock: 1,
			Price: 1000,
		}

		err := product.Validate()
		assert.Nil(t, err)
	})
	t.Run("product required", func(t *testing.T) {
		product := Product{
			Name:  "",
			Stock: 1,
			Price: 1000,
		}

		err := product.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("product name invalid", func(t *testing.T) {
		product := Product{
			Name:  "Pr",
			Stock: 1,
			Price: 1000,
		}

		err := product.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrProductInvalid, err)
	})
	t.Run("product stock invalid", func(t *testing.T) {
		product := Product{
			Name:  "Product 1",
			Stock: 0,
			Price: 1000,
		}

		err := product.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrStockInvalid, err)
	})
	t.Run("product price invalid", func(t *testing.T) {
		product := Product{
			Name:  "Product 1",
			Stock: 1,
			Price: 0,
		}

		err := product.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPriceInvalid, err)
	})
}
