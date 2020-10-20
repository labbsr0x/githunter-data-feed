package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

type CommentsRequest struct {
	Ids []string
}

func (c *Controller) GetCommentsHandler(ctx *fiber.Ctx) {
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")

	commentsRequest := &CommentsRequest{}
	if err := ctx.BodyParser(commentsRequest); err != nil {
		logrus.Warn(err)
		ctx.Next(fiber.NewError(fiber.StatusNoContent, "PARSER BODY: error parser input body"))
		return
	}

	if commentsRequest.Ids == nil {
		logrus.Warn("No ids defined in JSON request")
		ctx.Next(fiber.NewError(fiber.StatusNoContent, "No ids defined in JSON request"))
		return
	}

	data, err := c.Contract.GetComments(commentsRequest.Ids, provider, accessToken)
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
