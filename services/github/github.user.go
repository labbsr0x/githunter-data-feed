package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type UserResponse struct {
	User userInfo `json:"user"`
}

type userInfo struct {
	Name                      string                  `json:"name"`
	Login                     string                  `json:"login"`
	AvatarUrl                 string                  `json:"avatarUrl"`
	ContributionsCollection   contributionsCollection `json:"contributionsCollection"`
	RepositoriesContributedTo count                   `json:"repositoriesContributedTo"`
	PullRequests              count                   `json:"pullRequests"`
	Issues                    count                   `json:"issues"`
	Followers                 count                   `json:"followers"`
	Repositories              repos                   `json:"repositories"`
}

type contributionsCollection struct {
	TotalCommits                 int `json:"totalCommitContributions"`
	RestrictedContributionsCount int `json:"restrictedContributionsCount"`
}

type repos struct {
	NodeRepo []stargazers `json:"nodes"`
}

type stargazers struct {
	Stargazers count `json:"stargazers"`
}

func GetUserStats(login string, accessToken string) (*UserResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)
	if err != nil {
		return nil, err
	}

	maxQuantityReposStars := env.Get().Counters.NumberOfMaxQuantityItens

	respData := &UserResponse{}
	variables := map[string]interface{}{
		"login":                 login,
		"maxQuantityReposStars": maxQuantityReposStars,
	}

	err = client.Query(
		`query userStats($login:String!, $maxQuantityReposStars:Int!) {
			user(login:$login) {
				name,
				login,
				avatarUrl,
				contributionsCollection {
					totalCommitContributions,
					restrictedContributionsCount
				},
				repositoriesContributedTo(first: 1, contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY]) {
					totalCount
				},
				pullRequests(first: 1) {
					totalCount
				},
				issues(first: 1, filterBy:{ createdBy: $login }) {
					totalCount
				},
				followers {
					totalCount
				},
				repositories(first: $maxQuantityReposStars, ownerAffiliations: OWNER, isFork: false, orderBy: {direction: DESC, field: STARGAZERS}) {
					nodes {
						stargazers {
							totalCount
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
