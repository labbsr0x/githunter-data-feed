package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (c *Controller) GetMembersHandler(ctx *fiber.Ctx) {
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")
	organization := ctx.Query("organization")

	data, err := c.Contract.GetMembers(organization, provider, accessToken)
	if err != nil {
		logrus.Warn("Error requesting provider")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting provider"))
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
