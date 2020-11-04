package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type UserResponse struct {
	User userInfo `json:"user"`
}

type userInfo struct {
	Name                      string                    `json:"name"`
	Login                     string                    `json:"login"`
	AvatarUrl                 string                    `json:"avatarUrl"`
	Company                   string                    `json:"company"`
	Organizations             organizations             `json:"organizations"`
	ContributionsCollection   contributionsCollection   `json:"contributionsCollection"`
	RepositoriesContributedTo repositoriesContributedTo `json:"repositoriesContributedTo"`
	PullRequests              count                     `json:"pullRequests"`
	Issues                    count                     `json:"issues"`
	Followers                 count                     `json:"followers"`
	Repositories              repos                     `json:"repositories"`
}

type contributionsCollection struct {
	TotalCommits                 int `json:"totalCommitContributions"`
	RestrictedContributionsCount int `json:"restrictedContributionsCount"`
}

type repositoriesContributedTo struct {
	PathOfRepositoriesContributed []pathRepoContributed `json:"nodes"`
	TotalCount                    int                   `json:"totalCount"`
}

type pathRepoContributed struct {
	Name  string `json:"name"`
	Owner owner  `json:"owner"`
}

type owner struct {
	Login string `json:"login"`
}

type organizations struct {
	Organization []user `json:"nodes"`
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
	maxQuantityReposContributed := env.Get().Counters.NumberOfMaxQuantityItens

	respData := &UserResponse{}
	variables := map[string]interface{}{
		"login":                       login,
		"maxQuantityReposStars":       maxQuantityReposStars,
		"maxQuantityReposContributed": maxQuantityReposContributed,
	}

	err = client.Query(
		`query userStats($login:String!, $maxQuantityReposStars:Int!, $maxQuantityReposContributed:Int!) {
			user(login:$login) {
				name,
				login,
				organizations(first:100){
					nodes{
						login
					}
				}
				company,
				avatarUrl,
				contributionsCollection {
					totalCommitContributions,
					restrictedContributionsCount
				},
				repositoriesContributedTo(first: $maxQuantityReposContributed, contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY]) {
					nodes {
            name, 
            owner {
              login
						}
					},
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
