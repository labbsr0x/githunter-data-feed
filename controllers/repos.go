package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-repos/github"
	"github.com/rs/zerolog/log"
)

// GetRepos function
func GetRepos(c *fiber.Ctx) {

	// param passed by param URL
	accessToken := c.Params("access_token")
	println(accessToken)

	repos := github.GetLastRepos(10, accessToken)
	if repos == nil {
		log.Error().Msg("Error requesting github")
		c.Next(fiber.NewError(fiber.StatusInternalServerError, "Error requesting github"))
		return
	}

	repositoreis := [10]string{}
	for k, v := range repos.Viewer.Repositories.Nodes {
		repositoreis[k] = v.Name
	}
	result := map[string]interface{}{
		"name":         repos.Viewer.Name,
		"repositories": repositoreis,
	}

	b, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Msg("error at unmarshal")
		return
	}

	c.Fasthttp.Response.Header.Add("Content-type", "application/json")
	c.Status(fiber.StatusOK).Send(b)
}
