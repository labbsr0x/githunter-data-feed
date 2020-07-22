package graphql

import (
	"context"
	"fmt"

	ql "github.com/machinebox/graphql"
)

// Graphql define the object wich will be created with funtion New
type Graphql struct {
	Query func(string, map[string]interface{}, interface{}) error
}

//New create a new GraphQL instance
func New(clientURL string, accessToken string) (*Graphql, error) {
	var err error

	graphql := &Graphql{}

	client := ql.NewClient(clientURL)

	if client == nil {
		err = fmt.Errorf("Couldn't possible to make new GraphQL Client")
	}

	graphql.Query = buildQuery(client, accessToken)

	return graphql, err
}

// buildQuery is a closure function to use the client and accessToken parameters
// passed in New function and create a new function wich will be used to
// mount and send a GraphQL request
func buildQuery(client *ql.Client, accessToken string) func(string, map[string]interface{}, interface{}) error {
	return func(query string, variables map[string]interface{}, resp interface{}) error {
		req := ql.NewRequest(query)

		//todo: implements message
		if req == nil {
			return fmt.Errorf("Couldn't possible to make new GraphQL Client")
		}

		for key, value := range variables {
			req.Var(key, value)
		}

		auth(req, accessToken)

		//todo: implement logRUS
		if err := client.Run(context.Background(), req, &resp); err != nil {
			return err
		}

		return nil
	}
}

// Add authorization header in request
func auth(req *ql.Request, accessToken string) {
	req.Header.Add("Authorization", "Bearer "+accessToken)
}
