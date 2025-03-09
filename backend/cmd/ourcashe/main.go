package main

import (
	//"finance-api/handlers"
	"finance-api/internal/config"
	"finance-api/internal/database"
	"finance-api/internal/handlers"
	"finance-api/internal/middleware"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config - cleanenv
	cfg := config.MustLoad()

	// init logger - slog
	log := setupLogger(cfg.Env)
	log.Info("start", slog.String("env", cfg.Env))
	log.Debug("Debug enabled")

	// init databse - mysql
	database.InitDb(cfg.User, cfg.Password, cfg.Host, cfg.Name, cfg.Port)

	app := fiber.New()
	// JWT Middleware
	apiV1 := app.Group("/api/v1")
	apiV1Transaction := apiV1.Group("/transaction")
	apiV1Transaction.Get("/get", middleware.JWTProtected(), handlers.HandlerGetTransactions)
	apiV1Transaction.Post("/add", middleware.JWTProtected(), handlers.HandlerAddTransaction)
	apiV1Transaction.Delete("/delete/:id", middleware.JWTProtected(), handlers.HandlerDeleteTransaction)
	apiV1Auth := apiV1.Group("/auth")
	apiV1Auth.Post("/signup", handlers.HandlerSignup)
	apiV1Auth.Post("/signin", handlers.HandlerSignin)
	apiV1Auth.Post("/logout", middleware.JWTProtected(), handlers.HandlerLogout)
	apiV1Auth.Get("/user", middleware.JWTProtected(), handlers.HandlerUser)

	app.Listen(":8080")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
