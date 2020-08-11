package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (c *Controller) GetPullsHandler(ctx *fiber.Ctx) {
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")
	owner := ctx.Query("owner")
	repo := ctx.Query("name")

	issues, err := c.Contract.GetPulls(10, owner, repo, provider, accessToken)
	if err != nil {
		logrus.Warn("Error requesting github")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	b, err := json.Marshal(issues)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}

	ctx.Fasthttp.Response.Header.Add("Content-type", "application/json")
	ctx.Status(fiber.StatusOK).Send(b)
}
