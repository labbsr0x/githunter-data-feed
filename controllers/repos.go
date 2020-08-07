package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

// GetRepos function
func (c *Controller) GetReposHandler(ctx *fiber.Ctx) {

	// param passed by param URL
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")

	repos, err := c.Contract.GetLastRepos(10, accessToken, provider)
	if err != nil {
		logrus.Warn("Error requesting github")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	b, err := json.Marshal(repos)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}

	ctx.Fasthttp.Response.Header.Add("Content-type", "application/json")
	ctx.Status(fiber.StatusOK).Send(b)
}
