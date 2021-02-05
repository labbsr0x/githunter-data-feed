package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
)

type UserScoreResponse struct {
	User userScoreInfo `json:"user"`
}

type userScoreInfo struct {
	Name                      string                    `json:"name"`
	Login                     string                    `json:"login"`
	ID                        string                    `json:"id"`
	AvatarUrl                 string                    `json:"avatarUrl"`
	Followers                 followers                 `json:"followers"`
	Organizations             organizations             `json:"organizations"`
	Company                   string                    `json:"company"`
	RepositoriesContributedTo repositoriesContributedTo `json:"repositoriesContributedTo"`
	ContributionsCollection   contributionsCollection   `json:"contributionsCollection"`
	Repositories              repos                     `json:"repositories"`
}

func GetUserScore(login string, accessToken string) (*UserScoreResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	respData := &UserScoreResponse{}

	numberCount := env.Get().Counters.NumberOfLastItens
	numberQuantity := env.Get().Counters.NumberOfQuantityItens
	maxQuantityFollowers := env.Get().Counters.NumberOfMaxQuantityItens
	maxQuantityOrganizations := env.Get().Counters.NumberOfMaxQuantityItens
	maxQuantityReposStars := env.Get().Counters.NumberOfMaxQuantityItens
	maxQuantityReposContributed := env.Get().Counters.NumberOfMaxQuantityItens
	maxQuantityPullRequests := env.Get().Counters.NumberOfMaxQuantityItens
	maxQuantityIssues := env.Get().Counters.NumberOfMaxQuantityItens

	variables := map[string]interface{}{
		"login":                       login,
		"numberCount":                 numberCount,
		"numberQuantity":              numberQuantity,
		"maxQuantityFollowers":        maxQuantityFollowers,
		"maxQuantityOrganizations":    maxQuantityOrganizations,
		"maxQuantityPullRequests":     maxQuantityPullRequests,
		"maxQuantityIssues":           maxQuantityIssues,
		"maxQuantityReposStars":       maxQuantityReposStars,
		"maxQuantityReposContributed": maxQuantityReposContributed,
	}

	err = client.Query(
		`query userScore($login:String!, $numberCount:Int!, $maxQuantityFollowers:Int!, $maxQuantityOrganizations:Int!, $numberQuantity:Int!, $maxQuantityReposStars:Int!, $maxQuantityReposContributed:Int!, $maxQuantityPullRequests:Int!, $maxQuantityIssues:Int!) {
			user(login: $login) {
				name
				login
				id
				avatarUrl
    			followers(first:$maxQuantityFollowers) {
					nodes {
						login
					}
				},
				organizations(first:$maxQuantityOrganizations) {
					nodes {
						login
					}
				},
				company,
				repositoriesContributedTo(last: $maxQuantityReposContributed, contributionTypes: [COMMIT]){
					nodes{
						name
						owner{
							login
						}
					}
				},
				contributionsCollection{
					pullRequestContributions(last: $maxQuantityPullRequests) {
						totalCount
						nodes{
							pullRequest{
								number
								repository{
									name
									owner {
										login
									}
								}
								state
								createdAt
								updatedAt
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
								comments (first: $numberQuantity) {
									totalCount
									nodes {
									id
									url
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
					},
					issueContributions(last: $maxQuantityIssues) {
						totalCount
						nodes{
							issue{
								number
								repository{
									name
									owner {
										login
									}
								}
								state
								createdAt
								updatedAt
								closedAt
								author {
									login
								}
								labels(first: $numberCount, orderBy: {field: CREATED_AT, direction: DESC}) {
									nodes {
										name
									}
								}
								participants(last: $numberCount) {
									totalCount
									nodes {
										 login
										}
								}
								timelineItems(last: $numberQuantity, itemTypes: ISSUE_COMMENT) {
									totalCount
									updatedAt
									nodes {
		
										__typename
										... on IssueComment {
											id
											url
											createdAt
											author {
												login
											}
										}
									}
								}
							}
						}
					}
				}
				repositories(last: $maxQuantityReposStars, ownerAffiliations: OWNER, isFork: false, orderBy: {direction: DESC, field: STARGAZERS}) {
					nodes {
						name
						owner {
							login
						}
						createdAt
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
