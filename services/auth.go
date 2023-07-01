package services

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mrcharlss/3thingstoday/dal"
	"github.com/mrcharlss/3thingstoday/types"
	"github.com/mrcharlss/3thingstoday/utils"
	"gorm.io/gorm"
)

func Signup(ctx *fiber.Ctx) error{
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

    err := dal.FindUserByEmail(&struct{ID string}{},b.Email ).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

    fmt.Printf("%v", b)

    user := &dal.User{
        Name: b.Name,
        Password: "pass",
        Email: b.Email,
    }

    if err := dal.CreateUser(user); err.Error != nil {
        return fiber.NewError(fiber.StatusConflict, err.Error.Error())
    }

    return nil
}
