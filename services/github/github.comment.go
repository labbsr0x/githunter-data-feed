package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
)

type CommentsResponse struct {
	Nodes []commentNode `json:"nodes"`
}

type commentNode struct {
	Repository commentRepository `json:"repository"`
	CreatedAt  string            `json:"createdAt"`
	Issue      commentIssue      `json:"issue"`
	Url        string            `json:"url"`
	Id         string            `json:"id"`
	Author     user              `json:"author"`
}

type commentRepository struct {
	Name  string `json:"name"`
	Owner user   `json:"owner"`
}

type commentIssue struct {
	Number int `json:"number"`
}

func GetComments(ids []string, accessToken string) (*CommentsResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)
	if err != nil {
		return nil, err
	}

	respData := &CommentsResponse{}
	variables := map[string]interface{}{
		"ids": ids,
	}

	err = client.Query(
		`query ($ids: [ID!]!) {
			  nodes(ids: $ids) {
				__typename
				... on IssueComment {
				  repository {
					name
					owner {
					  login
					}
				  }
				  createdAt
				  issue {
					number
				  }
				  url
				  id
				  author {
					login
				  }
				}
			  }
			}`, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
