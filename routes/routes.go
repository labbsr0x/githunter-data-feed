package routes

import (
	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/controllers"
)

//Register routes
func Register(app *fiber.App, version string) {

	v1 := app.Group("/" + version)

	//TODO: JWT //:email?:provider
	v1.Get("/repos", func(c *fiber.Ctx) {
		controllers.GetRepos(c)
	})

}
