package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/sirupsen/logrus"
)

type ReposController struct {
	Contract services.Contract
}

func NewReposController() *ReposController {
	theController := &ReposController{
		Contract: services.New(),
	}
	return theController
}

// GetRepos function
func (r *ReposController) GetReposHandler(c *fiber.Ctx) {

	// param passed by param URL
	accessToken := c.Query("access_token")
	provider := c.Query("provider")

	repos, err := r.Contract.GetLastRepos(10, accessToken, provider)
	if err != nil {
		logrus.Warn("Error requesting github")
		c.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	b, err := json.Marshal(repos)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}

	c.Fasthttp.Response.Header.Add("Content-type", "application/json")
	c.Status(fiber.StatusOK).Send(b)
}
