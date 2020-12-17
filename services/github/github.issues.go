package github

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/graphql"
)

type issue struct {
	TotalCount int         `json:"totalCount"`
	Nodes      []issueNode `json:"nodes"`
}

type issueNode struct {
	Number        int          `json:"number"`
	Repository    pathRepo     `json:"repository"`
	State         string       `json:"state"`
	CreatedAt     string       `json:"createdAt"`
	UpdatedAt     string       `json:"updatedAt"`
	ClosedAt      string       `json:"closedAt"`
	Author        user         `json:"author"`
	Labels        labels       `json:"labels"`
	Participants  participants `json:"participants"`
	TimelineItems comments     `json:"timelineItems"`
}

func GetIssues(owner string, repo string, accessToken string) (*Response, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	respData := &Response{}

	query := `query($number_of_issues:Int!, $owner:String!, $repo:String!) {
							repository(name: $repo, owner: $owner) {
								issues(last: $number_of_issues) {
									nodes {
										number
										state
										createdAt
										updatedAt
										closedAt
										author {
											login
										}
										labels(first: 10, orderBy: {field: CREATED_AT, direction: DESC}) {
											nodes {
												name
											}
										}
										participants(last: 10) {
											totalCount
											nodes {
												 login
											  }
										}
										timelineItems(last: 10, itemTypes: ISSUE_COMMENT) {
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
						}`

	numberOfIssues := env.Get().Counters.NumberOfLastItens

	variables := map[string]interface{}{
		"number_of_issues": numberOfIssues,
		"owner":            owner,
		"repo":             repo,
	}

	err = client.Query(query, variables, respData)

	if err != nil {
		return nil, err
	}

	return respData, nil
}
