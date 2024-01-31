package auth

import (
	"log"
	"testing"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dicki@mail.com",
			Password: "rahasiasecret",
		}

		err := authEntity.Validate()
		assert.Nil(t, err)
	})
	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "rahasiasecret",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrEmailRequired, err)
	})
	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dickimail.com",
			Password: "rahasiasecret",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrEmailInvalid, err)
	})
	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dicki@mail.com",
			Password: "",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPasswordRequired, err)
	})
	t.Run("password is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dicki@mail.com",
			Password: "root",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPasswordInvalid, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dicki@mail.com",
			Password: "rahasiasecret",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)

		assert.Nil(t, err)
		log.Printf("%+v\n", authEntity)
	})
}
