package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type CommitsResponse struct {
	Viewer commitsViewer `json:"repository"`
}

type commitsViewer struct {
	DefaultBranch commitsDefaultBranch `json:"defaultBranchRef"`
}

type commitsDefaultBranch struct {
	Target commitsTarget `json:"target"`
}

type commitsTarget struct {
	Commits commitsNode `json:"commits"`
}

type commitsNode struct {
	Nodes []commit `json:"nodes"`
}

type commit struct {
	Message       string `json:"message"`
	CommittedDate string `json:"committedDate"`
}

func GetCommitsRepo(nameRepo string, ownerRepo string, accessToken string) (*CommitsResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)
	if err != nil {
		return nil, err
	}

	respData := &CommitsResponse{}
	variables := map[string]interface{}{
		"name":     nameRepo,
		"owner":    ownerRepo,
		"quantity": 10,
	}

	err = client.Query(
		`query getInfoCommitsPage($name:String!, $owner:String!, $quantity:Int!) {
			repository(name: $name, owner: $owner) {
				...RepoFragmentCommits,
			}
		}
		
		fragment RepoFragmentCommits on Repository {
			defaultBranchRef {
				target {
					... on Commit {
						commits: history(first: $quantity) { 
							nodes {
								message,
								committedDate,
							}
						}
					}
				}
			}
		}`, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
