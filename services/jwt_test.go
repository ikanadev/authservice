package services_test

import (
	"testing"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/matryer/is"
	"github.com/vmkevv/authservice/services"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Should generate a token", func(t *testing.T) {
		is := is.New(t)
		_, err := services.GenerateToken("secret", 1, 1, time.Millisecond)
		is.NoErr(err) // Error should be nil
	})
	t.Run("Token should be invalid if its expired", func(t *testing.T) {
		is := is.New(t)
		token, _ := services.GenerateToken("secret", 1, 1, time.Millisecond)
		time.Sleep(1000 * time.Millisecond)
		_, _, err := services.CheckToken("secret", token)
		is.Equal(err, jwt.ErrExpValidation)
	})
}
