package main

import (
	"github.com/Ahmad940/dropify/app/handler/ws"
	"github.com/Ahmad940/dropify/pkg/middleware"
	"github.com/Ahmad940/dropify/pkg/router"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var cfg = websocket.Config{
	RecoverHandler: func(conn *websocket.Conn) {
		if err := recover(); err != nil {
			conn.WriteJSON(fiber.Map{"customError": "error occurred"})
		}
	},
}

func AttachRoutes(app *fiber.App) {

	base := app.Group("/")

	app.Get("/ws/:id", websocket.New(ws.Stream, cfg))

	api := base.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"message": "Hello, World!"}) })

	// routes
	router.Authentication(api)

	// not found
	middleware.NotFoundMiddleware(app)
}
