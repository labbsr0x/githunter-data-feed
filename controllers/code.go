package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/sirupsen/logrus"
)

type CodeController struct {
	Contract services.Contract
}

func NewCodeController() *CodeController {
	theController := &CodeController{
		Contract: services.New(),
	}
	return theController
}

// GetCode function
func (c *CodeController) GetCodeHandler(ctx *fiber.Ctx) {

	// param passed by param URL
	// name and owner
	name := ctx.Query("name")
	owner := ctx.Query("owner")
	accessToken := ctx.Query("access_token")
	provider := ctx.Query("provider")
	//check value query!!

	data, err := c.Contract.GetInfoCodePage(name, owner, accessToken, provider)
	if err != nil {
		logrus.Warn("Error requesting github")
		ctx.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	if data == nil {
		logrus.Warn("data response is null")
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
