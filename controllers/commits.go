package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/sirupsen/logrus"
)

type CommitsController struct {
	Contract services.Contract
}

func NewCommitsController() *CommitsController {
	theController := &CommitsController{
		Contract: services.New(),
	}
	return theController
}

// GetRepos function
func (r *CommitsController) GetCommitsHandler(c *fiber.Ctx) {

	// param passed by param URL
	name := c.Query("name")
	owner := c.Query("owner")
	quantity, err := strconv.Atoi(c.Query("quantity"))
	accessToken := c.Query("access_token")
	provider := c.Query("provider")

	commits, err := r.Contract.GetCommitsRepo(name, owner, quantity, accessToken, provider)
	if err != nil {
		logrus.Warn("Error requesting github")
		c.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	b, err := json.Marshal(commits)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("error at unmarshal")
		return
	}

	c.Fasthttp.Response.Header.Add("Content-type", "application/json")
	c.Status(fiber.StatusOK).Send(b)
}
