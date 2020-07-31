package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/labbsr0x/githunter-api/services/gitlab"
	"github.com/sirupsen/logrus"
)

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
}

type defaultContract struct{}


func New() Contract {
	return &defaultContract{}
}

// ReposResponseContract
type ReposResponseContract struct {
	Name         string   `json:"name"`
	Repositories []string `json:"repositories"`
}

func (d *defaultContract) GetLastRepos(numberOfRepos int, accessToken string, provider string) (*ReposResponseContract, error ){

	theContract := &ReposResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetLastRepos(numberOfRepos, accessToken)
		break
	case `gitlab`:
		theContract, err = gitlabGetLastRepos(numberOfRepos, accessToken)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetLastRepos unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetLastRepos returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetLastRepos(numberOfRepos int, accessToken string) (*ReposResponseContract, error) {

	repos, err := github.GetLastRepos(numberOfRepos, accessToken)

	if err != nil {
		return nil, err
	}

	repositories := []string{}
	for _, v := range repos.Viewer.Repositories.Nodes {
		repositories = append(repositories, v.Name)
	}
	result := &ReposResponseContract{
		Name:         repos.Viewer.Name,
		Repositories: repositories,
	}

	return result, nil
}

func gitlabGetLastRepos(numberOfRepos int, accessToken string) (*ReposResponseContract, error) {

	repos, err := gitlab.GetLastRepos(numberOfRepos, accessToken)

	if err != nil {
		return nil, err
	}

	repositoreis := []string{}
	for _, v := range repos.Projects.Nodes {
		repositoreis = append(repositoreis, v.Name)
	}
	result := &ReposResponseContract{
		Name:         repos.CurrentUser.Name,
		Repositories: repositoreis,
	}

	return result, nil
}
