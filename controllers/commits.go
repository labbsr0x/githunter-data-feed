package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

// GetRepos function
func (c *Controller) GetCommitsHandler(ctx *fiber.Ctx) {
	// param passed by param URL
	name := ctx.Query("name")
	owner := ctx.Query("owner")
	quantity, err := strconv.Atoi(ctx.Query("quantity"))
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")

	data, err := c.Contract.GetCommitsRepo(name, owner, quantity, accessToken, provider)
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
