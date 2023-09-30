package main

import (
	"embed"
	"net/http"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

// Embed a directory
//
//go:embed dist/*
var embedDirStatic embed.FS

func EnableMiddlewares(app *fiber.App) {
	// middleware
	app.Use(cors.New())
	app.Use(recover.New())

	app.Use(logger.New())

	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Route().Path, "/ws")
		},
	}))

	// http middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "dist",
	}))

	app.Get("/monitor", monitor.New())

}
