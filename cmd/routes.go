package main

import (
	"github.com/gioeleber/go-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/fact", handlers.CreateFactPage)
	app.Post("/fact", handlers.CreateFact)
}