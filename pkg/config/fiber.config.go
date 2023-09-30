package config

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		// UnescapePath: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
}
