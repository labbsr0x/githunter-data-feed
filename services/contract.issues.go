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
	Number       int          `json:"number"`
	State        string       `json:"state"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
	ClosedAt     string       `json:"closedAt"`
	Author       string       `json:"author"`
	Labels       []string     `json:"labels"`
	Comments     comments     `json:"comments"`
	Participants participants `json:"participants"`
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
	issuesResp, err := github.GetIssues(numberOfIssues, owner, repo, accessToken)

	if err != nil {
		return nil, err
	}

	issues := []issue{}
	for _, v := range issuesResp.Repository.Issues.Nodes {
		theIssue := issue{}
		theIssue.Number = v.Number
		theIssue.State = v.State
		theIssue.Author = v.Author.Login
		theIssue.CreatedAt = v.CreatedAt
		theIssue.UpdatedAt = v.UpdatedAt
		theIssue.ClosedAt = v.ClosedAt

		theIssue.Participants.TotalCount = v.Participants.TotalCount
		for _, t := range v.Participants.User {
			theIssue.Participants.User = append(theIssue.Participants.User, t.Login)
		}

		for _, l := range v.Labels.Label {
			theIssue.Labels = append(theIssue.Labels, l.Name)
		}

		theIssue.Comments.TotalCount = v.TimelineItems.TotalCount

		for _, t := range v.TimelineItems.Data {
			theComment := comment{}
			theComment.Author = t.Author.Login
			theComment.CreatedAt = t.CreatedAt
			theIssue.Comments.Data = append(theIssue.Comments.Data, theComment)
		}

		issues = append(issues, theIssue)
	}

	result := &IssuesResponseContract{
		Issues: issues,
	}

	return result, nil
}
