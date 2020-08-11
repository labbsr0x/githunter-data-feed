package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type pulls struct {
	TotalCount int        `json:"totalCount"`
	Nodes      []pullNode `json:"nodes"`
}

type pullNode struct {
	Number       int          `json:"number"`
	State        string       `json:"state"`
	CreatedAt    string       `json:"createdAt"`
	Merged       bool         `json:"merged"`
	MergedAt     string       `json:"mergedAt"`
	ClosedAt     string       `json:"closedAt"`
	Author       user         `json:"author"`
	Labels       labels       `json:"labels"`
	Participants participants `json:"participants"`
	Comments     comments     `json:"comments"`
}

func GetPulls(numberCount int, owner string, name string, accessToken string, closed bool) (*Response, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	states := []string{"OPEN"}
	if closed == true {
		states = []string{"CLOSED", "MERGED"}
	}

	respData := &Response{}

	query := `query($numberCount:Int!, $owner:String!, $name:String!, $states:[PullRequestState!]) {
				  repository(name: $name, owner: $owner) {
					pullRequests(last: $numberCount, states: $states) {
					  totalCount
					  nodes{
						number
						state
						createdAt
						closedAt
						merged
						mergedAt
						author {
							login
						}
						labels (first: $numberCount) {
						  nodes {
							name
						  }
						}
						comments (first: $numberCount) {
						  totalCount
						  nodes {
							createdAt
							author {
							  login
							}
						  }
						}
						participants (last: $numberCount){
						  totalCount
						  nodes {
							 login
						  }
						}
					  }
					}
				  }
				}`

	variables := map[string]interface{}{
		"numberCount": numberCount,
		"owner":       owner,
		"name":        name,
		"state":       states,
	}

	err = client.Query(query, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
