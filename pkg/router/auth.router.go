package router

import (
	"github.com/Ahmad940/dropify/app/handler"
	"github.com/gofiber/fiber/v2"
)

func Authentication(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
