package github

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/services/graphql"
	"github.com/sirupsen/logrus"
)

type IssuesResponse struct {
	Repository repository `json:"repository"`
}

type repository struct {
	Issues issue `json:"issues"`
}

type issue struct {
	Nodes []issueNode `json:"nodes"`
}

type issueNode struct {
	Number        int          `json:"number"`
	State         string       `json:"state"`
	CreatedAt     string       `json:"createdAt"`
	UpdatedAt     string       `json:"updatedAt"`
	ClosedAt      string       `json:"closedAt"`
	Author        author       `json:"author"`
	Labels        label        `json:"labels"`
	Participants  participant  `json:"participants"`
	TimelineItems timelineItem `json:"timelineItems"`
}

type author struct {
	Login string `json:"login"`
}

type label struct {
	Nodes []labelNode `json:"nodes"`
}

type labelNode struct {
	Name string `json:"name"`
}

type participant struct {
	TotalCount int `json:"totalCount"`
}

type timelineItem struct {
	TotalCount int                `json:"totalCount"`
	UpdatedAt  string             `json:"updatedAt"`
	Nodes      []timelineItemNode `json:"nodes"`
}

type timelineItemNode struct {
	TypeName  string `json:"__typename"`
	CreatedAt string `json:"createdAt"`
	Author    author `json:"author"`
}

func GetIssues(numberOfIssues int, owner string, repo string, accessToken string) (*IssuesResponse, error) {
	client, err := graphql.New(env.Get().GithubGraphQLURL, accessToken)

	if err != nil {
		return nil, err
	}

	respData := &IssuesResponse{}

	variables := map[string]interface{}{
		"number_of_issues": numberOfIssues,
		"owner":            owner,
		"repo":             repo,
	}

	logrus.Info("Start query")

	err = client.Query(`query($number_of_issues:Int!, $owner:String!, $repo:String!) {
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
		}`, variables, respData)

	logrus.Info("End query")
	if err != nil {
		return nil, err
	}

	return respData, nil
}
