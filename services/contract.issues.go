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
	Number            int           `json:"number"`
	State             string        `json:"state"`
	CreatedAt         string        `json:"createdAt"`
	UpdatedAt         string        `json:"updatedAt"`
	ClosedAt          string        `json:"closedAt"`
	Author            string        `json:"author"`
	Labels            []string      `json:"labels"`
	TotalParticipants int           `json:"totalParticipants"`
	TimelineItems     timelineItems `json:"timelineItems"`
}

type timelineItems struct {
	TotalCount int            `json:"totalCount"`
	UpdatedAt  string         `json:"updatedAt"`
	Items      []timelineItem `json:"nodes"`
}

type timelineItem struct {
	TypeName  string `json:"__typename"`
	CreatedAt string `json:"createdAt"`
	Author    string `json:"author"`
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
	issues, err := github.GetIssues(numberOfIssues, owner, repo, accessToken)

	if err != nil {
		return nil, err
	}

	issuesTemp := []issue{}
	issueTemp := issue{}
	timelineItemTemp := timelineItem{}

	for _, v := range issues.Repository.Issues.Nodes {
		issueTemp.Number = v.Number
		issueTemp.State = v.State
		issueTemp.Author = v.Author.Login
		issueTemp.CreatedAt = v.CreatedAt
		issueTemp.UpdatedAt = v.UpdatedAt
		issueTemp.ClosedAt = v.ClosedAt
		issueTemp.TotalParticipants = v.Participants.TotalCount

		for _, l := range v.Labels.Nodes {
			issueTemp.Labels = append(issueTemp.Labels, l.Name)
		}

		issueTemp.TimelineItems.TotalCount = v.TimelineItems.TotalCount
		issueTemp.TimelineItems.UpdatedAt = v.TimelineItems.UpdatedAt

		for _, t := range v.TimelineItems.Nodes {
			timelineItemTemp.Author = t.Author.Login
			timelineItemTemp.CreatedAt = t.CreatedAt
			timelineItemTemp.TypeName = t.TypeName

			issueTemp.TimelineItems.Items = append(issueTemp.TimelineItems.Items, timelineItemTemp)
		}

		issuesTemp = append(issuesTemp, issueTemp)
	}

	result := &IssuesResponseContract{
		Issues: issuesTemp,
	}

	return result, nil
}
