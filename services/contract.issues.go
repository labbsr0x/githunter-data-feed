package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

type IssuesResponseContract struct {
	Issues []issue `json:"issues"`
}

type issue struct {
	Number        int           `json:"number"`
	State         string        `json:"state"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	ClosedAt      string        `json:"closedAt"`
	Author        author        `json:"author"`
	Labels        []string      `json:"labels"`
	Participants  participant   `json:"participants"`
	TimelineItems timelineItems `json:"timelineItems"`
}

type author struct {
	Login string `json:"login"`
}

type participant struct {
	TotalCount int `json:"totalCount"`
}

type timelineItems struct {
	TotalCount int            `json:"totalCount"`
	UpdatedAt  string         `json:"updatedAt"`
	Items      []timelineItem `json:"nodes"`
}

type timelineItem struct {
	TypeName  string `json:"__typename"`
	CreatedAt string `json:"createdAt"`
	Author    author `json:"author"`
}

func (d *defaultContract) GetIssues(numberOfIssues int, owner string, repo string, provider string, accessToken string) (*IssuesResponseContract, error) {
	theContract := &IssuesResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetIssues(numberOfIssues, owner, repo, accessToken)
		break
	case `gitlab`:
		//theContract, err = gitlabGetIssues(numberOfIssues, owner, repo, accessToken)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetIssues unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetIssues returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetIssues(numberOfIssues int, owner string, repo string, accessToken string) (*IssuesResponseContract, error) {
	logrus.Info("Start")
	issues, err := github.GetIssues(numberOfIssues, owner, repo, accessToken)
	logrus.Info("End")

	if err != nil {
		return nil, err
	}
	logrus.Info("Start 2")
	issueTemp := issue{}
	issuesResp := []issue{}
	for _, v := range issues.Repository.Issues.Nodes {
		issueTemp.Number = v.Number
		issueTemp.State = v.State
		issueTemp.Author.Login = v.Author.Login
		issueTemp.CreatedAt = v.CreatedAt
		issueTemp.UpdatedAt = v.UpdatedAt
		issueTemp.ClosedAt = v.ClosedAt
		issuesResp = append(issuesResp, issueTemp)
	}
	logrus.Info("End 2")
	result := &IssuesResponseContract{
		Issues: issuesResp,
	}
	return result, nil
}
