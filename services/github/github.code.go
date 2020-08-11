package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type CodeResponse struct {
	Viewer codeViewer `json:"repository"`
}

type codeViewer struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	CreatedAt        string            `json:"createdAt"`
	PrimaryLanguage  node              `json:"primaryLanguage"`
	RepositoryTopics repositoryTopics  `json:"repositoryTopics"`
	Watchers         count             `json:"watchers"`
	Stars            count             `json:"stargazers"`
	Forks            int               `json:"forkCount"`
	LastCommit       nodeDefaultBranch `json:"defaultBranchRef"`
	Readme           text              `json:"readme"`
	Contributing     text              `json:"contributing"`
	LicenseInfo      node              `json:"licenseInfo"`
	CodeOfConduct    nodeCodeOfConduct `json:"codeOfConduct"`
	Releases         count             `json:"releases"`
	Contributors     history           `json:"totalContributors"`
	Languages        nodeLanguages     `json:"languages"`
	DiskUsage        int               `json:"diskUsage"`
}

type nodeDefaultBranch struct {
	DefaultBranch nodeTarget `json:"target"`
}

type nodeTarget struct {
	LastCommitDate string `json:"lastCommitDate"`
	CommitsQuanity count  `json:"commits"`
}

type nodeCodeOfConduct struct {
	Body         string `json:"body"`
	ResourcePath string `json:"resourcePath"`
}

type history struct {
	History count `json:"history"`
}

type nodeLanguages struct {
	Quantity  int             `json:"totalCount"`
	Languages []edgeLanguages `json:"edges"`
}

type edgeLanguages struct {
	Size     int      `json:"size"`
	Language language `json:"node"`
}

func GetInfoCodePage(nameRepo string, ownerRepo string, accessToken string) (*CodeResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)
	if err != nil {
		return nil, err
	}

	respData := &CodeResponse{}
	variables := map[string]interface{}{
		"name":  nameRepo,
		"owner": ownerRepo,
		"count": 100,
		"zero":  0,
	}
	err = client.Query(
		`query getInfoCodePage($name:String!, $owner:String!, $count:Int!, $zero:Int!) {
			repository(name: $name, owner: $owner) {
				name,
				description,
				createdAt,
				primaryLanguage { name },
				repositoryTopics(first: $count) { 
					nodes {
						topic {
							name
						}
					}
				},
				watchers { totalCount },
				stargazers { totalCount },
				forkCount,
				...RepoFragmentCommits,
				readme: object(expression: "master:README.md") {
					... on Blob { text }
				},
				contributing: object(expression: "master:CONTRIBUTING.md") {
					... on Blob { text }
				},
				licenseInfo { name },
				codeOfConduct {
					body,
					resourcePath,
				},
				releases(first: $zero) {
					totalCount
				},
				#fundingLinks { url },
				totalContributors: object(expression:"master") {
					... on Commit {
						history { totalCount }
					}
				},
				languages(first: $count, orderBy: { field: SIZE, direction: DESC }) {
					totalCount,
					edges { 
						size, 
						node { language: name } 
					}
				},
				diskUsage
			}
		}
		
		fragment RepoFragmentCommits on Repository {
			defaultBranchRef {
				target {
					... on Commit {
						lastCommitDate: committedDate,
						commits: history(first: $zero) { totalCount }
					}
				}
			}
		}`, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
