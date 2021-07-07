package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zeimedee/shortUrl/database"
	"github.com/zeimedee/shortUrl/handlers"
)

func Handler(app *fiber.App) {
	app.Get("/short/:short", handlers.Redirect)
	app.Post("/shortUrl", handlers.ShortenUrl)

}

func main() {
	database.ConnecDb()
	app := fiber.New()

	app.Use(cors.New())
	Handler(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // 404 not found
	})

	log.Fatal(app.Listen(":3000"))
}
