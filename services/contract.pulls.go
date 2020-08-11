package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

type PullsResponseContract struct {
	Total int `json:"total"`
	Pulls []pull `json:"pulls"`
}

type pull struct {
	Number            int           `json:"number"`
	State             string        `json:"state"`
	CreatedAt         string        `json:"createdAt"`
	ClosedAt          string        `json:"closedAt"`
	Merged            bool        	`json:"merged"`
	MergedAt          string        `json:"mergedAt"`
	Author            string        `json:"author"`
	Labels            []string      `json:"labels"`
	Comments     comments 		`json:"comments"`
	Participants participants `json:"participants"`
}

func (d *defaultContract) GetPulls(numberCount int, owner string, name string, provider string, accessToken string) (*PullsResponseContract, error) {
	theContract := &PullsResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetPulls(numberCount, owner, name, accessToken)
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

func githubGetPulls(numberCount int, owner string, repo string, accessToken string) (*PullsResponseContract, error) {
	pullsOpened, err := github.GetPulls(numberCount, owner, repo, accessToken, false)

	if err != nil {
		return nil, err
	}

	data := formatContract(pullsOpened)

	pullsClosed, err := github.GetPulls(numberCount, owner, repo, accessToken, true)

	if err != nil {
		return nil, err
	}

	data = append(data, formatContract(pullsClosed)...)

	result := &PullsResponseContract{
		Total: pullsOpened.Repository.Pulls.TotalCount + pullsClosed.Repository.Pulls.TotalCount,
		Pulls: data,
	}


	return result, nil
}

func formatContract(response *github.IssuesResponse) []pull {

	data := []pull{}
	for _, v := range response.Repository.Pulls.Nodes {
		theData := pull{}
		theData.Number = v.Number
		theData.State = v.State
		theData.Author = v.Author.Login
		theData.CreatedAt = v.CreatedAt
		theData.ClosedAt = v.ClosedAt
		theData.Merged = v.Merged
		theData.MergedAt = v.MergedAt

		theData.Participants.TotalCount = v.Participants.TotalCount
		for _, t := range v.Participants.User {
			theData.Participants.User = append(theData.Participants.User, t.Login)
		}

		for _, l := range v.Labels.Label {
			theData.Labels = append(theData.Labels, l.Name)
		}

		theData.Comments.TotalCount = v.Comments.TotalCount
		for _, t := range v.Comments.Data {
			theComment := comment{}
			theComment.Author = t.Author.Login
			theComment.CreatedAt = t.CreatedAt
			theData.Comments.Data = append(theData.Comments.Data, theComment)
		}

		data = append(data, theData)
	}

	return data
}