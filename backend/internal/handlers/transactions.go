package handlers

import (
	"finance-api/internal/database"
	"finance-api/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

var transactions []model.Transaction

func HandlerGetTransactions(c *fiber.Ctx) error {
	database.GetAllRecords(&transactions)
	return c.Status(fiber.StatusOK).JSON(transactions)
}

func HandlerAddTransaction(c *fiber.Ctx) error {
	var newTransaction model.Transaction
	if err := c.BodyParser(&newTransaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if newTransaction.Date.IsZero() {
		newTransaction.Date = time.Now()
	}
	err := database.InsertRecord(&newTransaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert transaction"})
	}
	return c.Status(fiber.StatusCreated).JSON(newTransaction)
}

func HandlerDeleteTransaction(c *fiber.Ctx) error {
	// Get the ID from request params
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transaction ID is required"})
	}

	// Attempt to delete the transaction
	err := database.DeleteRecordByID(&model.Transaction{}, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete transaction"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Transaction deleted successfully"})
}
