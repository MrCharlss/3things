package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrcharlss/3thingstoday/services"
)

func Things(app *fiber.App)  {
    r := app.Group("/things")

    r.Post("/create", services.CreateThings)
    r.Get("/list", services.GetThings)
    
}
