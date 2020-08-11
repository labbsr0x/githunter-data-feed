package routes

import (
	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/controllers"
)

//Register routes
func Register(app *fiber.App, version string) {

	v1 := app.Group("/" + version)

	theController := controllers.NewController()

	//TODO: JWT //:email?:provider
	v1.Get("/repos", theController.GetReposHandler)
	v1.Get("/issues", theController.GetIssuesHandler)
	v1.Get("/pulls", theController.GetPullsHandler)
}
