package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) {
	c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.SendString("New Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.SendString("Delete Book")
}
