package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Err represents a server error
type Err struct {
	HTTPCode int    `json:"-"`
	ErrCode  string `json:"code"`
	Message  string `json:"message"`
}

// Error to implement golang error interface
func (s Err) Error() string {
	return s.Message
}

// NewErr creates a new Err
func NewErr(errCode, message string, HTTPCode int) Err {
	return Err{HTTPCode, errCode, message}
}

// New500 creates a new Err with 500 Internal server error code
func New500(err error) Err {
	return Err{
		HTTPCode: fiber.StatusInternalServerError,
		Message:  fmt.Sprintf("Error %v", err.Error()),
	}
}

// FiberErrorHandler custom error handler for fiber
func FiberErrorHandler(c *fiber.Ctx, e error) error {
	err, ok := e.(Err)
	if ok {
		c.Status(err.HTTPCode)
		return c.JSON(err)
	}
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}
