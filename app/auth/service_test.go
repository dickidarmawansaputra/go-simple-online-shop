package auth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/external/database"
	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
	"github.com/google/uuid"
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

var email = fmt.Sprintf("%v@mail.com", uuid.NewString())

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := RegisterRequestPayload{
			Email:    email,
			Password: "rahasiasecret",
		}

		err := svc.register(context.Background(), req)
		assert.Nil(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		req := RegisterRequestPayload{
			Email:    email,
			Password: "rahasiasecret",
		}

		err := svc.register(context.Background(), req)
		log.Println("ERROR LOGIN FAIL")
		log.Println(err)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := LoginRequestPayload{
			Email:    email,
			Password: "rahasiasecret",
		}

		token, err := svc.login(context.Background(), req)
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
	t.Run("fail", func(t *testing.T) {
		req := LoginRequestPayload{
			Email:    email,
			Password: "rahasiasecretfail",
		}

		token, err := svc.login(context.Background(), req)
		assert.NotNil(t, err)
		assert.Empty(t, token)
	})
}
