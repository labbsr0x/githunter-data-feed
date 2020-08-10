package routes

import (
	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/controllers"
)

//Register routes
func Register(app *fiber.App, version string) {

	v1 := app.Group("/" + version)

	reposController := controllers.NewReposController()
	codeController := controllers.NewCodeController()

	//TODO: JWT //:email?:provider
	v1.Get("/repos", reposController.GetReposHandler)
	v1.Get("/code", codeController.GetCodeHandler)

}
