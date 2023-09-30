package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(app *fiber.App) {
	// custom 404
	app.Use(func(ctx *fiber.Ctx) error {
		routeUrl := ctx.OriginalURL()
		method := ctx.Method()
		// fmt.Printf("Value :%v\n", routeUrl)
		return ctx.Status(fiber.StatusNotFound).Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":    fmt.Sprintf("Cannot %v %v", method, routeUrl),
			"error":      "Not Found",
			"statusCode": fiber.StatusNotFound,
		})
	})
}
