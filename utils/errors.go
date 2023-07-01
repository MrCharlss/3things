package utils

import "github.com/gofiber/fiber/v2"

type httpError struct {
    StausCode int `json:"statusCode"`
    Error string `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError

    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }

    return c.Status(code).JSON(&httpError{
        StausCode: code,
        Error: err.Error(),
    })
}
