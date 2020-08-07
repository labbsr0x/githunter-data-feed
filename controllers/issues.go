package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (c *Controller) GetIssuesHandler(ctx *fiber.Ctx) {
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")
	owner := ctx.Query("owner")
	repo := ctx.Query("repo")

	issues, err := c.Contract.GetIssues(10, owner, repo, provider, accessToken)
	if err != nil {
		logrus.Warn("Error requesting github")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	logrus.Infof("Start Marshal")
	b, err := json.Marshal(issues)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}
	logrus.Infof("End Marshal")

	ctx.Fasthttp.Response.Header.Add("Content-type", "application/json")
	ctx.Status(fiber.StatusOK).Send(b)
}
