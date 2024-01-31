package transaction

import (
	"context"
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/external/database"
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

func TestCreateTransaction(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU: "9ccc727b-90b4-48d6-855b-a7ba5b795b27",
			Amount:     2,
			Email:      "user@user.com",
		}

		err := svc.CreateTransaction(context.Background(), req)
		assert.Nil(t, err)
	})
}
