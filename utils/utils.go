package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return Validate(body)

}

func GetUser(c *fiber.Ctx) *uint {
    id := uint(1)
	// id, _ := c.Locals("USER").(uint)
	return &id
}
