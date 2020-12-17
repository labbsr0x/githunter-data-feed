package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
)

type organization struct {
	MembersWithRole membersWithRole `json:"membersWithRole"`
}

type membersWithRole struct {
	Edges []memberEdge `json:"edges"`
}

type memberEdge struct {
	Cursor string     `json:"cursor"`
	Node   memberNode `json:"node"`
}

type memberNode struct {
	Login string `json:"login"`
}

func GetMembers(organization string, after string, accessToken string) (*Response, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	respData := &Response{}

	query := `query($organization:String!, $numberQuantity:Int!, $after:String) {
				  organization(login: $organization) { 
					membersWithRole(first:$numberQuantity, after: $after) {
					  edges {
						cursor
						node {
						  login
						}
					  }
					}
				  }
				}`

	numberQuantity := env.Get().Counters.NumberOfMaxQuantityItens

	var cursor *string
	cursor = &after
	if after == "" {
		cursor = nil
	}

	variables := map[string]interface{}{
		"organization":   organization,
		"numberQuantity": numberQuantity,
		"after":          cursor,
	}

	err = client.Query(query, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
