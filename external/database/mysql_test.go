package database

import (
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
	"github.com/stretchr/testify/assert"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionMysql(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectMysql(config.Cfg.DB)
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
