package github

import (
	"context"

	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
)

var (
	client *graphql.Client
)

// Response Fail
type ResponseFail struct {
	message          string `json:"message"`
	documentationURL string `json:"documentation_url"`
}

// Repos Response
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

type node struct {
	Name string `json:"name"`
}

func connect() {
	if client == nil {
		client = graphql.NewClient(env.Get().GithubGraphQLURL)
	}
}

// Get Last Repos
func GetLastRepos(numberOfRepos int, accessToken string) *ReposResponse {

	// Connect with Github if not.
	connect()

	req := graphql.NewRequest(`query($number_of_repos:Int!) {
		viewer {
		  name
		   repositories(last: $number_of_repos) {
			 nodes {
			   name
			 }
		   }
		 }
	  }`)

	req.Var("number_of_repos", numberOfRepos)

	auth(req, accessToken)

	respData := &ReposResponse{}
	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Warn().Msg(err.Error())
		return nil
	}

	return respData
}

func auth(req *graphql.Request, accessToken string) {
	req.Header.Add("Authorization", "bearer "+accessToken)
}
