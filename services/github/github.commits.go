package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
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
	Oid           string `json:"oid"`
	CommittedDate string `json:"committedDate"`
	Author        author `json:"author"`
}

func GetCommitsRepo(nameRepo string, ownerRepo string, authorID string, accessToken string) (*CommitsResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)
	if err != nil {
		return nil, err
	}

	quantity := env.Get().Counters.NumberOfLastItens

	respData := &CommitsResponse{}

	variables := map[string]interface{}{
		"name":     nameRepo,
		"owner":    ownerRepo,
		"quantity": quantity,
		"authorID": authorID,
	}

	var query string

	if authorID != "" {
		query = `query getInfoCommitsPage($name:String!, $owner:String!, $quantity:Int!, $authorID:ID!) {
			repository(name: $name, owner: $owner) {
				...RepoFragmentCommits,
			}
		}
		
		fragment RepoFragmentCommits on Repository {
			defaultBranchRef {
				target {
					... on Commit {
						commits: history(first: $quantity, author: {id: $authorID}) { 
							nodes {
								oid,
								message,
								committedDate,
								author{
									user {
										login
										id
									}
								}
							}
						}
					}
				}
			}
		}`
	} else {
		query = `query getInfoCommitsPage($name:String!, $owner:String!, $quantity:Int!) {
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
								oid,
								message,
								committedDate,
								author{
									user {
									login
									}
								}
							}
						}
					}
				}
			}
		}`
	}

	err = client.Query(query, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
