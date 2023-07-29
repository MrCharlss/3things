package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mrcharlss/3thingstoday/dal"
	"github.com/mrcharlss/3thingstoday/routes"
	"github.com/mrcharlss/3thingstoday/utils"
)

func main() {
	fmt.Println("Works")
	dal.Connect()
    dal.Migrate(&dal.User{}, &dal.Thing{})
	app := fiber.New(fiber.Config{
        ErrorHandler: utils.ErrorHandler,
    }) 
	app.Use(logger.New())

    routes.AuthRoutes(app)
    routes.Things(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello world")
	})

    app.Listen(":8000")
}
