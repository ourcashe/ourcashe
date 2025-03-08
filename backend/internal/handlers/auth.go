package handlers

import (
	"finance-api/internal/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandlerSignup(c *fiber.Ctx) error {
	payload := new(model.User)

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	log.Println("Username:", payload.Username)
	log.Println("Email:", payload.Email)
	log.Println("Password:", payload.Password)

	return c.SendString("Username: " + payload.Username + " Email: " + payload.Email + " Password: " + payload.Password)
}
