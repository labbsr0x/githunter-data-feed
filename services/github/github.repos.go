package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
)

// ResponseFail define the struct of a fail request
type ResponseFail struct {
	message          string `json:"message"`
	documentationURL string `json:"documentation_url"`
}

//ReposResponse define the return's struct of GetLastRepos
type ReposResponse struct {
	Viewer viewer `json:"viewer"`
}

type viewer struct {
	Name         string       `json:"name"`
	Repositories repositories `json:"repositories"`
}

type repositories struct {
	Nodes []node `json:"nodes"`
}

// GetLastRepos is used for get last repos of the user
func GetLastRepos(accessToken string) (*ReposResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	numberOfRepos := env.Get().Counters.NumberOfLastItens

	respData := &ReposResponse{}
	variables := map[string]interface{}{
		"number_of_repos": numberOfRepos,
	}

	err = client.Query(`query($number_of_repos:Int!) {
		viewer {
		  name
		   repositories(last: $number_of_repos) {
			 nodes {
			   name
			 }
		   }
		 }
		}`, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
