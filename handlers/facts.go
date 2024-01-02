package handlers

import (
	"github.com/gioeleber/go-trivia/models"
	"github.com/gioeleber/go-trivia/database"
	"github.com/gofiber/fiber/v2"
)


func Home(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title": "Go Trivia | Home",
		"Facts": facts,
	})
}

func CreateFactPage(c *fiber.Ctx) error {
	return c.Render("create-fact", fiber.Map{
		"Title": "Go Trivia | create fact",
	})
}


func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}