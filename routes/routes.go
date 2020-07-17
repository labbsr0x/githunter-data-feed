package routes

import (
	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-repos/controllers"
)

//Register routes
func Register(app *fiber.App, version string) {

	v1 := app.Group("/" + version)

	v1.Get("/repos/:access_token", func(c *fiber.Ctx) {
		controllers.GetRepos(c)
	})

}
