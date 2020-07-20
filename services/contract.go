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

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract = githubGetLastRepos(numberOfRepos, accessToken)
		break
	case `gitlab`:
		theContract = gitlabGetLastRepos(numberOfRepos, accessToken)
		break
	}

	if(theContract == nil){
		logrus..Debug("GetLastRepos returned a null answer")
	}

	return theContract
}

func githubGetLastRepos(numberOfRepos int, accessToken string) *ReposResponseContract {

	repos := github.GetLastRepos(numberOfRepos, accessToken)

	if repos == nil {
		return nil
	}

	repositoreis := []string{}
	for _, v := range repos.Viewer.Repositories.Nodes {
		repositoreis = append(repositoreis, v.Name)
	}
	result := &ReposResponseContract{
		Name:         repos.Viewer.Name,
		Repositories: repositoreis,
	}

	return result
}

func gitlabGetLastRepos(numberOfRepos int, accessToken string) *ReposResponseContract {

	repos := gitlab.GetLastRepos(numberOfRepos, accessToken)

	if repos == nil {
		return nil
	}

	repositoreis := []string{}
	for _, v := range repos.Projects.Nodes {
		repositoreis = append(repositoreis, v.Name)
	}
	result := &ReposResponseContract{
		Name:         repos.CurrentUser.Name,
		Repositories: repositoreis,
	}

	return result
}
