package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

// GetUserStats function
func (c *Controller) GetUserScoreHandler(ctx *fiber.Ctx) {
	login := ctx.Query("login")
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")

	data, err := c.Contract.GetUserScore(login, accessToken, provider)
	if err != nil {
		logrus.Warn("Error requesting github")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	if data == nil {
		logrus.Warn("No data found")
		ctx.Next(fiber.NewError(fiber.StatusNoContent, "No data found"))
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}

	ctx.Fasthttp.Response.Header.Add("Content-type", "application/json")
	ctx.Status(fiber.StatusOK).Send(b)
}
