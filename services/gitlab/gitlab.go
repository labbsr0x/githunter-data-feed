package gitlab

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
	CurrentUser node     `json:"currentUser"`
	Projects    projects `json:"projects"`
}

type projects struct {
	Nodes []node `json:"nodes"`
}

type node struct {
	Name string `json:"name"`
}

func connect() {
	client = graphql.NewClient(env.Get().GitlabGraphQLURL)
}

// Get Last Repos
func GetLastRepos(numberOfRepos int, accessToken string) *ReposResponse {

	// Connect with Github if not.
	connect()

	req := graphql.NewRequest(`query($number_of_repos:Int!) {
			currentUser {
			  name
			}
			projects(last:$number_of_repos membership: true) {
			  nodes {
				name
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
	req.Header.Add("Authorization", "Bearer "+accessToken)
}
