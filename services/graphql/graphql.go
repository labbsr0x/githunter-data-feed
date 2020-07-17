package graphql

import (
	"context"

	ql "github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
)

// Graphql define the object wich will be created with funtion New
type Graphql struct {
	Query func(string, map[string]interface{}, interface{})
}

//New create a new GraphQL instance
func New(clientURL string, accessToken string) *Graphql {
	graphql := &Graphql{}

	client := ql.NewClient(clientURL)

	graphql.Query = buildQuery(client, accessToken)

	return graphql
}

// buildQuery is a closure function to use the client and accessToken parameters
// passed in New function and create a new function wich will be used to
// mount and send a GraphQL request
func buildQuery(client *ql.Client, accessToken string) func(string, map[string]interface{}, interface{}) {
	return func(query string, variables map[string]interface{}, resp interface{}) {
		req := ql.NewRequest(query)

		for key, value := range variables {
			req.Var(key, value)
		}

		auth(req, accessToken)

		//todo: implement logRUS
		if err := client.Run(context.Background(), req, &resp); err != nil {
			log.Warn().Msg(err.Error())
		}
	}
}

// Add authorization header in request
func auth(req *ql.Request, accessToken string) {
	req.Header.Add("Authorization", "Bearer "+accessToken)
}
