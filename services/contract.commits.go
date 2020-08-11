package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

//struct responde of Commits page
type CommitsResponseContract struct {
	Commits []commit `json:"commits"`
}

type commit struct {
	Message       string `json:"message"`
	CommittedDate string `json:"committedDate"`
}

func (d *defaultContract) GetCommitsRepo(nameRepo string, ownerRepo string, quantity int, accessToken string, provider string) (*CommitsResponseContract, error) {
	theContract := &CommitsResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetCommitsRepo(nameRepo, ownerRepo, quantity, accessToken)
		break
	case `gitlab`:
		// theContract, err = githubGetCommitsRepo(nameRepo, ownerRepo, quantity, accessToken)
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

func githubGetCommitsRepo(nameRepo string, ownerRepo string, quantity int, accessToken string) (*CommitsResponseContract, error) {
	commits, err := github.GetCommitsRepo(nameRepo, ownerRepo, quantity, accessToken)
	if err != nil {
		return nil, err
	}

	commitsInfo := []commit{}
	for _, commitInf := range commits.Viewer.DefaultBranch.Target.Commits.Nodes {
		commitsInfo = append(commitsInfo, commit{
			Message:       commitInf.Message,
			CommittedDate: commitInf.CommittedDate,
		})
	}

	result := &CommitsResponseContract{
		Commits: commitsInfo,
	}

	return result, nil
}

func gitlabGetCommitsRepo(nameRepo string, ownerRepo string, quantity int, accessToken string) (*CommitsResponseContract, error) {

	return nil, nil
}
