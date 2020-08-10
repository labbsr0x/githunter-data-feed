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
func (r *CodeController) GetCodeHandler(c *fiber.Ctx) {

	// param passed by param URL
	// name and owner
	name := c.Query("name")
	owner := c.Query("owner")
	accessToken := c.Query("access_token")
	provider := c.Query("provider")
	//check value query!!

	repos, err := r.Contract.GetInfoCodePage(name, owner, accessToken, provider)
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
