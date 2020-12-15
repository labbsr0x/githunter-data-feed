package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

// struct reponse of User score
type UserScoreResponseContract struct {
	Name                    string                 `json:"name"`
	Login                   string                 `json:"login"`
	ID                      string                 `json:"id"`
	ContributedRepositories []nameOwnerContributed `json:"contributedRepositories"`
	Pulls                   []pull                 `json:"pulls"`
	Issues                  []issue                `json:"issues"`
	OwnedRepositories       []ownedRepository      `json:"ownedRepositories"`
}

type ownedRepository struct {
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	CreatedAt     string `json:"createdAt"`
	StarsReceived int    `json:"starsReceived"`
}

func (d *defaultContract) GetUserScore(login string, accessToken string, provider string) (*UserScoreResponseContract, error) {

	theContract := &UserScoreResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetUserScore(login, accessToken)
		break
	case `gitlab`:
		gitlabClient = gitlabNewClient(accessToken)
		theContract, err = gitlabGetUserScore(login)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetUserScore unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetUserScore returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetUserScore(login string, accessToken string) (*UserScoreResponseContract, error) {
	data, err := github.GetUserScore(login, accessToken)

	if err != nil {
		return nil, err
	}

	nameOwnercontributions := []nameOwnerContributed{}
	for _, r := range data.User.RepositoriesContributedTo.PathOfRepositoriesContributed {
		nameOwnercontributions = append(nameOwnercontributions, nameOwnerContributed{
			Owner: r.Owner.Login,
			Name:  r.Name,
		})
	}

	issues := []issue{}
	for _, v := range data.User.ContributionsCollection.IssueContributions.Nodes {
		theIssue := issue{}
		theIssue.Number = v.Issue.Number
		theIssue.Name = v.Issue.Repository.Name
		theIssue.Owner = v.Issue.Repository.Owner.Login
		theIssue.State = v.Issue.State
		theIssue.Author = v.Issue.Author.Login
		theIssue.CreatedAt = v.Issue.CreatedAt
		theIssue.UpdatedAt = v.Issue.UpdatedAt
		theIssue.ClosedAt = v.Issue.ClosedAt

		theIssue.Participants.TotalCount = v.Issue.Participants.TotalCount
		for _, t := range v.Issue.Participants.User {
			theIssue.Participants.User = append(theIssue.Participants.User, t.Login)
		}

		for _, l := range v.Issue.Labels.Label {
			theIssue.Labels = append(theIssue.Labels, l.Name)
		}

		theIssue.Comments.TotalCount = v.Issue.TimelineItems.TotalCount

		for _, t := range v.Issue.TimelineItems.Data {
			theComment := shortComment{}
			theComment.ID = t.ID
			theComment.URL = t.URL
			theComment.Author = t.Author.Login
			theComment.CreatedAt = t.CreatedAt
			theIssue.Comments.Data = append(theIssue.Comments.Data, theComment)
		}

		issues = append(issues, theIssue)
	}

	pulls := []pull{}
	for _, v := range data.User.ContributionsCollection.PullRequestContributions.Nodes {
		theData := pull{}
		theData.Number = v.Pull.Number
		theData.Name = v.Pull.Repository.Name
		theData.Owner = v.Pull.Repository.Owner.Login
		theData.State = v.Pull.State
		theData.Author = v.Pull.Author.Login
		theData.CreatedAt = v.Pull.CreatedAt
		theData.UpdatedAt = v.Pull.UpdatedAt
		theData.ClosedAt = v.Pull.ClosedAt
		theData.Merged = v.Pull.Merged
		theData.MergedAt = v.Pull.MergedAt

		theData.Participants.TotalCount = v.Pull.Participants.TotalCount
		for _, t := range v.Pull.Participants.User {
			theData.Participants.User = append(theData.Participants.User, t.Login)
		}

		for _, l := range v.Pull.Labels.Label {
			theData.Labels = append(theData.Labels, l.Name)
		}

		theData.Comments.TotalCount = v.Pull.Comments.TotalCount
		for _, t := range v.Pull.Comments.Data {
			theComment := shortComment{}
			theComment.ID = t.ID
			theComment.URL = t.URL
			theComment.Author = t.Author.Login
			theComment.CreatedAt = t.CreatedAt
			theData.Comments.Data = append(theData.Comments.Data, theComment)
		}

		pulls = append(pulls, theData)
	}

	ownedRepositories := []ownedRepository{}
	for _, r := range data.User.Repositories.NodeRepo {
		ownedRepositories = append(ownedRepositories, ownedRepository{
			Owner:         r.Owner.Login,
			Name:          r.Name,
			CreatedAt:     r.CreatedAt,
			StarsReceived: r.Stargazers.TotalCount,
		})
	}

	result := &UserScoreResponseContract{
		Name:                    data.User.Name,
		Login:                   data.User.Login,
		ID:                      data.User.ID,
		ContributedRepositories: nameOwnercontributions,
		Pulls:                   pulls,
		Issues:                  issues,
		OwnedRepositories:       ownedRepositories,
	}

	return result, nil
}

func gitlabGetUserScore(login string) (*UserScoreResponseContract, error) {

	result := &UserScoreResponseContract{}

	return result, nil
}
