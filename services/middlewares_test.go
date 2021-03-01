package services_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/matryer/is"
	"github.com/vmkevv/authservice/services"
)

func TestWithAuthMiddleware(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: services.FiberErrorHandler,
	})
	app.Use(services.WithAuth([]int8{1}, "secret"))
	app.Get("/id", func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		return c.JSON(userID)
	})
	app.Get("/role", func(c *fiber.Ctx) error {
		role := c.Locals("role")
		return c.JSON(role)
	})

	t.Run("Should set userID as context value", func(t *testing.T) {
		is := is.New(t)
		token, _ := services.GenerateToken("secret", 1, 1, 1*time.Minute)
		req, _ := http.NewRequest("GET", "/id", nil)
		req.Header.Set("Authorization", token)

		resp, _ := app.Test(req)
		is.Equal(fiber.StatusOK, resp.StatusCode)
		respString, _ := ioutil.ReadAll(resp.Body)
		is.Equal("1", string(respString))
	})

	t.Run("Should set role as context value", func(t *testing.T) {
		is := is.New(t)
		token, _ := services.GenerateToken("secret", 1, 1, 1*time.Minute)
		req, _ := http.NewRequest("GET", "/role", nil)
		req.Header.Set("Authorization", token)

		resp, _ := app.Test(req)
		is.Equal(fiber.StatusOK, resp.StatusCode)
		respString, _ := ioutil.ReadAll(resp.Body)
		is.Equal("1", string(respString))
	})

	t.Run("Should return 401 with the wrong role", func(t *testing.T) {
		is := is.New(t)
		token, _ := services.GenerateToken("secret", 1, 2, 1*time.Minute)
		req, _ := http.NewRequest("GET", "/role", nil)
		req.Header.Set("Authorization", token)

		resp, _ := app.Test(req)
		is.Equal(fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Should return 401 when there is no token", func(t *testing.T) {
		is := is.New(t)
		req, _ := http.NewRequest("GET", "/role", nil)
		resp, _ := app.Test(req)
		is.Equal(fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Should return 401 when there is a random token", func(t *testing.T) {
		is := is.New(t)
		req, _ := http.NewRequest("GET", "/role", nil)
		req.Header.Set("Authorization", "randomString")
		resp, _ := app.Test(req)
		is.Equal(fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Should return correct message when token is expired", func(t *testing.T) {
		is := is.New(t)
		token, _ := services.GenerateToken("secret", 1, 1, 1*time.Microsecond)
		req, _ := http.NewRequest("GET", "/role", nil)
		req.Header.Set("Authorization", token)

		time.Sleep(1 * time.Second)
		resp, _ := app.Test(req)
		respString, _ := ioutil.ReadAll(resp.Body)
		expectedErr := services.NewErr(services.ExpiredToken, "JWT token expired", fiber.StatusUnauthorized)
		expectedStr, _ := json.Marshal(expectedErr)
		is.Equal(string(expectedStr), string(respString))
	})
}
