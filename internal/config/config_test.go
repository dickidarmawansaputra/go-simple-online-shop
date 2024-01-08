package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		filename := "../../cmd/api/config.yaml"
		err := LoadConfig(filename)

		assert.Nil(t, err)
		log.Printf("%+v\n", Cfg)
	})
}
