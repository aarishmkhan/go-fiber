package main

import (
	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) {
	c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
