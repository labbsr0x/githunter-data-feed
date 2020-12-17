package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-data-feed/services/github"
	"github.com/sirupsen/logrus"
)

// struct reponse of User stats
type UserResponseContract struct {
	Name                    string                 `json:"name"`
	Login                   string                 `json:"login"`
	AvatarUrl               string                 `json:"avatarUrl"`
	Company                 string                 `json:"company"`
	Organizations           []string               `json:"organizations"`
	ContributedRepositories []nameOwnerContributed `json:"contributedRepositories"`
	Amount                  amount                 `json:"amount"`
}

type amount struct {
	TotalReposContributed int `json:"totalReposContributed"`
	Commits               int `json:"commits"`
	PullRequests          int `json:"pullRequests"`
	Issues                int `json:"issuesOpened"`
	StarsReceived         int `json:"starsReceived"`
	Followers             int `json:"followers"`
}

type nameOwnerContributed struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
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
		return nil, fmt.Errorf("GetUserInfo unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetUserInfo returned a null answer")
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

	organizations := []string{}
	for _, u := range data.User.Organizations.Organization {
		organizations = append(organizations, u.Login)
	}

	nameOwnercontributions := []nameOwnerContributed{}
	for _, r := range data.User.RepositoriesContributedTo.PathOfRepositoriesContributed {
		nameOwnercontributions = append(nameOwnercontributions, nameOwnerContributed{
			Owner: r.Owner.Login,
			Name:  r.Name,
		})
	}

	amount := &amount{
		TotalReposContributed: data.User.RepositoriesContributedTo.TotalCount,
		Commits:               (data.User.ContributionsCollection.TotalCommits + data.User.ContributionsCollection.RestrictedContributionsCount),
		PullRequests:          data.User.PullRequests.TotalCount,
		Issues:                data.User.Issues.TotalCount,
		StarsReceived:         totalStarsReceived,
		Followers:             data.User.Followers.TotalCount,
	}

	result := &UserResponseContract{
		Name:                    data.User.Name,
		Login:                   data.User.Login,
		Company:                 data.User.Company,
		Organizations:           organizations,
		AvatarUrl:               data.User.AvatarUrl,
		ContributedRepositories: nameOwnercontributions,
		Amount:                  *amount,
	}

	return result, nil
}

func gitlabGetUserStats(login string) (*UserResponseContract, error) {

	result := &UserResponseContract{}

	return result, nil
}
