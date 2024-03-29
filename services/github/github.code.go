package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
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
	HomepageUrl      string            `json:" homepageUrl"`
	Readme           byteSize          `json:"readme"`
	Contributing     byteSize          `json:"contributing"`
	LicenseInfo      node              `json:"licenseInfo"`
	CodeOfConduct    nodeCodeOfConduct `json:"codeOfConduct"`
	Releases         count             `json:"releases"`
	// Contributors     history           `json:"totalContributors"` *Info no longer available*
	Languages nodeLanguages `json:"languages"`
	DiskUsage int           `json:"diskUsage"`
}

type nodeDefaultBranch struct {
	DefaultBranch nodeTarget `json:"target"`
}

type nodeTarget struct {
	LastCommitDate string `json:"lastCommitDate"`
	CommitsQuanity count  `json:"commits"`
}

type nodeCodeOfConduct struct {
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

	maxQuantityTopics := env.Get().Counters.NumberOfLastItens
	maxQuantityLangs := env.Get().Counters.NumberOfQuantityItens

	respData := &CodeResponse{}
	variables := map[string]interface{}{
		"name":              nameRepo,
		"owner":             ownerRepo,
		"maxQuantityTopics": maxQuantityTopics,
		"maxQuantityLangs":  maxQuantityLangs,
		"zero":              0,
	}
	err = client.Query(
		`query getInfoCodePage($name:String!, $owner:String!, $maxQuantityTopics:Int!, $maxQuantityLangs:Int!, $zero:Int!) {
			repository(name: $name, owner: $owner) {
				name,
				description,
				createdAt,
				primaryLanguage { name },
				repositoryTopics(first: $maxQuantityTopics) { 
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
				homepageUrl,
				readme: object(expression: "master:README.md") {
					... on Blob { byteSize }
				},
				contributing: object(expression: "master:CONTRIBUTING.md") {
					... on Blob { byteSize }
				},
				licenseInfo { name },
				codeOfConduct {
					resourcePath,
				},
				releases(first: $zero) {
					totalCount
				},
				languages(first: $maxQuantityLangs, orderBy: { field: SIZE, direction: DESC }) {
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
