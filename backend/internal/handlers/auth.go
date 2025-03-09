package handlers

import (
	"fmt"
	"strconv"
	"time"

	"finance-api/internal/database"
	"finance-api/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// JWT Secret Key (Load from env in real projects)

func HandlerSignup(c *fiber.Ctx) error {
	var payload model.User

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}
	user := model.User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err = database.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func HandlerSignin(c *fiber.Ctx) error {

	var payload model.LoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	var user model.User

	err := database.GetUserByEmail(&user, payload.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}
	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte("sadshnakjdhsa"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "could not login"})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	fmt.Print(token)
	return c.JSON(fiber.Map{"message": "success"})
}
func HandlerUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("sadshnakjdhsa"), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "unauthcicated"})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user model.User
	err = database.GetUserById(&user, claims.Issuer)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	return c.JSON(user)
}
func HandlerLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "Log out"})
}
