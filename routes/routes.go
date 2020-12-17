package routes

import (
	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-data-feed/controllers"
)

//Register routes
func Register(app *fiber.App, version string) {

	v1 := app.Group("/" + version)

	theController := controllers.NewController()

	//TODO: JWT //:email?:provider
	v1.Get("/repos", theController.GetReposHandler)
	v1.Get("/code", theController.GetCodeHandler)
	v1.Get("/commits", theController.GetCommitsHandler)
	v1.Get("/issues", theController.GetIssuesHandler)
	v1.Get("/pulls", theController.GetPullsHandler)
	v1.Get("/organization/members", theController.GetMembersHandler)
	v1.Get("/user/stats", theController.GetUserHandler)
	v1.Get("/userscore", theController.GetUserScoreHandler)
	v1.Post("/comments", theController.GetCommentsHandler)
}
