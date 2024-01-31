package transaction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSetSubTotal(t *testing.T) {
	var trx = Transaction{
		ProductPrice: 1000,
		Amount:       10,
	}

	trx.SetSubTotal()
	assert.Equal(t, uint(10000), trx.SubTotal)
}

func TestSetGrandTotal(t *testing.T) {
	t.Run("without platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 1000,
			Amount:       10,
		}

		trx.SetSubTotal()
		trx.SetGrandTotal()
		assert.Equal(t, uint(10000), trx.GrandTotal)
	})
	t.Run("with platform fee", func(t *testing.T) {
		var trx = Transaction{
			ProductPrice: 1000,
			Amount:       10,
			PlatformFee:  5000,
		}

		trx.SetSubTotal()
		trx.SetGrandTotal()
		assert.Equal(t, uint(15000), trx.GrandTotal)
	})
}

func TestProductJSON(t *testing.T) {
	product := Product{
		SKU:   uuid.NewString(),
		Name:  "Product JSON",
		Stock: 10,
		Price: 1000,
	}

	var trx = Transaction{}
	err := trx.SetProductJSON(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, trx.ProductJSON)

	productFromTrx, err := trx.GetProduct()
	assert.Nil(t, err)
	assert.NotEmpty(t, product, productFromTrx)
}

func TestTransactionStatus(t *testing.T) {
	type TableTest struct {
		Title    string
		Expected string
		Trx      Transaction
	}

	var tableTest = []TableTest{
		{
			Title:    "status created",
			Trx:      Transaction{Status: TransactionStatusCreated},
			Expected: TRX_CREATED,
		},
		{
			Title:    "status in progress",
			Trx:      Transaction{Status: TransactionStatusProgress},
			Expected: TRX_IN_PROGRESS,
		},
		{
			Title:    "status in delivery",
			Trx:      Transaction{Status: TransactionStatusInDelivery},
			Expected: TRX_IN_DELIVERY,
		},
		{
			Title:    "status completed",
			Trx:      Transaction{Status: TransactionStatusCompleted},
			Expected: TRX_IN_COMPLETED,
		},
		{
			Title:    "status unknown",
			Trx:      Transaction{Status: 0},
			Expected: TRX_UNKNOWN,
		},
	}

	for _, test := range tableTest {
		t.Run(test.Title, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Trx.GetStatus())
		})
	}
}
