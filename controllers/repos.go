package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-repos/services"
	"github.com/rs/zerolog/log"
)

// GetRepos function
func GetRepos(c *fiber.Ctx) {

	// param passed by param URL
	accessToken := c.Query("access_token")
	provider := c.Query("provider")

	repos := services.GetLastRepos(10, accessToken, provider)
	if repos == nil {
		log.Error().Msg("Error requesting github")
		c.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	b, err := json.Marshal(repos)
	if err != nil {
		log.Error().Err(err).Msg("error at unmarshal")
		return
	}

	c.Fasthttp.Response.Header.Add("Content-type", "application/json")
	c.Status(fiber.StatusOK).Send(b)
}
