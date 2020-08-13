package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

// GetCode function
func (c *Controller) GetCodeHandler(ctx *fiber.Ctx) {
	name := ctx.Query("name")
	owner := ctx.Query("owner")
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")

	data, err := c.Contract.GetInfoCodePage(10, 50, name, owner, accessToken, provider)
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
