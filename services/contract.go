package services

import (
	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/labbsr0x/githunter-api/services/gitlab"
	"github.com/sirupsen/logrus"
)

// ReposResponseContract
type ReposResponseContract struct {
	Name         string   `json:"name"`
	Repositories []string `json:"repositories"`
}

func GetLastRepos(numberOfRepos int, accessToken string, provider string) *ReposResponseContract {

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
	}

	if theContract == nil {
		logrus.Debug("GetLastRepos returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract
}

func githubGetLastRepos(numberOfRepos int, accessToken string) (*ReposResponseContract, error) {

	repos, err := github.GetLastRepos(numberOfRepos, accessToken)

	if err != nil {
		return nil, err
	}

	repositoreis := []string{}
	for _, v := range repos.Viewer.Repositories.Nodes {
		repositoreis = append(repositoreis, v.Name)
	}
	result := &ReposResponseContract{
		Name:         repos.Viewer.Name,
		Repositories: repositoreis,
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
