package main

import (
    "github.com/gioeleber/go-trivia/database"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "log"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}