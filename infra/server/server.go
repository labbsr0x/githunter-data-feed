package server

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/labbsr0x/githunter-repos/routes"
)

// Config is a function
func Config(version string) {

	app := fiber.New()

	app.Settings.Prefork = false
	app.Settings.CaseSensitive = true
	app.Settings.StrictRouting = true
	app.Settings.ServerHeader = "githunter"
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send()
	})
	app.Get("/error", func(c *fiber.Ctx) {
		c.Status(500)
		c.JSON(map[string]string{"message": "unknown error"})
	})
	routes.Register(app, version)

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
		c.Send("This is a dummy route")
	})
	app.Listen(3000)
}
