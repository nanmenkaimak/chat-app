package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nanmenkaimak/chat-app/internal/handlers"
	"log"
)

func routes() {
	router := fiber.New()
	router.Use(logger.New())

	auth := router.Group("/auth")
	{
		auth.Post("/signup", handlers.Repo.SignUp)
		auth.Post("/login", handlers.Repo.Login)
		auth.Post("/renew_access", handlers.Repo.RenewAccessToken)
	}

	log.Fatal(router.Listen(portNumber))
}
