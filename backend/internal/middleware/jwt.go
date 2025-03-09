package middleware

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Load secret from environment variable
var jwtSecret = []byte("sadshnakjdhsa")

// Middleware to check JWT from cookies
func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check for JWT token in Authorization header
		authHeader := c.Get("Authorization")

		// If Authorization header is missing, check cookies
		if authHeader == "" {
			token := c.Cookies("jwt")
			fmt.Println(token)
			if token == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": true,
					"msg":   "Missing or malformed JWT",
				})
			}
			// Set token in Authorization header so jwtware can process it
			c.Request().Header.Set("Authorization", "Bearer "+token)
		}

		// Pass request to jwtware middleware for validation
		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: jwtSecret},
			ContextKey: "jwt", // Store claims in Fiber context
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": true,
					"msg":   "Invalid or expired token",
				})
			},
		})(c)
	}
}
