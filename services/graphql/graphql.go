package graphql

import (
	"context"

	ql "github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

// Graphql define the object wich will be created with funtion New
type Graphql struct {
	Query func(string, map[string]interface{}, interface{})
}

//New create a new GraphQL instance
func New(clientURL string, accessToken string) *Graphql {
	graphql := &Graphql{}

	client := ql.NewClient(clientURL)

	if client == nil {
		logrus.WithFields(logrus.Fields{
			"client-url": clientURL,
		}).Warn("Couldn't possible to make new GraphQL Client")
	}

	graphql.Query = buildQuery(client, accessToken)

	return graphql
}

// buildQuery is a closure function to use the client and accessToken parameters
// passed in New function and create a new function wich will be used to
// mount and send a GraphQL request
func buildQuery(client *ql.Client, accessToken string) func(string, map[string]interface{}, interface{}) {
	return func(query string, variables map[string]interface{}, resp interface{}) {
		req := ql.NewRequest(query)

		if req == nil {
			logrus.WithFields(logrus.Fields{
				"query": query,
			}).Warn("Couldn't possible to make a new GraphQL request")
		}

		for key, value := range variables {
			req.Var(key, value)
		}

		auth(req, accessToken)

		//todo: implement logRUS
		if err := client.Run(context.Background(), req, &resp); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("An error ocurred while run a GraphQL request")
		}
	}
}

// Add authorization header in request
func auth(req *ql.Request, accessToken string) {
	req.Header.Add("Authorization", "Bearer "+accessToken)
}
