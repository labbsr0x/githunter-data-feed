package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
)

type issue struct {
	TotalCount int `json:"totalCount"`
	Nodes []issueNode `json:"nodes"`
}

type issueNode struct {
	Number        int          `json:"number"`
	State         string       `json:"state"`
	CreatedAt     string       `json:"createdAt"`
	UpdatedAt     string       `json:"updatedAt"`
	ClosedAt      string       `json:"closedAt"`
	Author        user       	`json:"author"`
	Labels        labels        `json:"labels"`
	Participants  participants  `json:"participants"`
	TimelineItems comments `json:"timelineItems"`
}

func GetIssues(numberOfIssues int, owner string, repo string, accessToken string) (*IssuesResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	respData := &IssuesResponse{}

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
