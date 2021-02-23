package authservice

import (
	"github.com/gofiber/fiber/v2"
)

// Err represents a server error
type Err struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// Error to implement golang error interface
func (s Err) Error() string {
	return s.Message
}

// NewErr creates a new Err
func NewErr(code int, message string) Err {
	return Err{code, message}
}

// FiberErrorHandler custom error handler for fiber
func FiberErrorHandler(c *fiber.Ctx, e error) error {
	err, ok := e.(Err)
	if ok {
		c.Status(err.Code)
		return c.JSON(err)
	}
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}
