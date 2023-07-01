package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrcharlss/3thingstoday/services"
)

func AuthRoutes(app fiber.Router)  {
    r := app.Group("/auth")

    r.Post("/signup", services.Signup)
    
}
