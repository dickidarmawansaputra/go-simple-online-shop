package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		tokenString, err := GenerateToken(publicId, "user", "secret")
		assert.Nil(t, err)
		assert.NotEmpty(t, tokenString)
		log.Println(tokenString)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		tokenString, err := GenerateToken(publicId, "user", "secret")
		assert.Nil(t, err)
		assert.NotEmpty(t, tokenString)

		jwtId, jwtRole, err := ValidateToken(tokenString, "secret")
		assert.Nil(t, err)
		assert.NotEmpty(t, jwtId)
		assert.NotEmpty(t, jwtRole)
		log.Println(tokenString)
	})
}
