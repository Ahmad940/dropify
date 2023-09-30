package middleware

import (
	"github.com/Ahmad940/dropify/pkg/config"
	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtWare.Config{
		SigningKey: jwtWare.SigningKey{
			Key: []byte(config.GetEnv().JWT_SECRET),
		},
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtWare.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err.Error() == "signature is invalid" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "token signature is invalid",
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
