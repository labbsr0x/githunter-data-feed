package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

// struct reponse of User stats
type UserResponseContract struct {
	Name   string `json:"name"`
	Login  string `json:"login"`
	Amount amount `json:"amount"`
}

type amount struct {
	RepositoryContribution int `json:"repositories"`
	Commits                int `json:"commits"`
	PullRequests           int `json:"pullRequests"`
	Issues                 int `json:"issues"`
	StarsReceived          int `json:"starsReceived"`
	Followers              int `json:"followers"`
}

func (d *defaultContract) GetUserStats(login string, accessToken string, provider string) (*UserResponseContract, error) {

	theContract := &UserResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetUserStats(login, accessToken)
		break
	case `gitlab`:
		gitlabClient = gitlabNewClient(accessToken)
		theContract, err = gitlabGetUserStats(login)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetUserStats unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetUserStats returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetUserStats(login string, accessToken string) (*UserResponseContract, error) {
	data, err := github.GetUserStats(login, accessToken)
	if err != nil {
		return nil, err
	}

	totalStarsReceived := 0
	for _, repo := range data.User.Repositories.NodeRepo {
		totalStarsReceived = +repo.Stargazers.TotalCount
	}

	amount := &amount{
		RepositoryContribution: data.User.RepositoriesContributedTo.TotalCount,
		Commits:                (data.User.ContributionsCollection.TotalCommits + data.User.ContributionsCollection.RestrictedContributionsCount),
		PullRequests:           data.User.PullRequests.TotalCount,
		Issues:                 data.User.Issues.TotalCount,
		StarsReceived:          totalStarsReceived,
		Followers:              data.User.Followers.TotalCount,
	}

	result := &UserResponseContract{
		Name:   data.User.Name,
		Login:  data.User.Login,
		Amount: *amount,
	}

	return result, nil
}

func gitlabGetUserStats(login string) (*UserResponseContract, error) {

	result := &UserResponseContract{}

	return result, nil
}
