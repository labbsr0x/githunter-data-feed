package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/services/github"
	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

//struct responde of Commits page
type CommitsResponseContract struct {
	Commits []commit `json:"data"`
}

type commit struct {
	Message       string `json:"message"`
	Oid           string `json:"number"`
	CommittedDate string `json:"committedDate"`
	Author        string `json:"author"`
}

func (d *defaultContract) GetCommitsRepo(nameRepo string, ownerRepo string, accessToken string, provider string) (*CommitsResponseContract, error) {
	theContract := &CommitsResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetCommitsRepo(nameRepo, ownerRepo, accessToken)
		break
	case `gitlab`:
		gitlabClient = gitlabNewClient(accessToken)
		theContract, err = gitlabGetCommits(nameRepo, ownerRepo, accessToken)
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

func githubGetCommitsRepo(nameRepo string, ownerRepo string, accessToken string) (*CommitsResponseContract, error) {
	commits, err := github.GetCommitsRepo(nameRepo, ownerRepo, accessToken)
	if err != nil {
		return nil, err
	}

	commitsInfo := []commit{}
	for _, commitInf := range commits.Viewer.DefaultBranch.Target.Commits.Nodes {
		commitsInfo = append(commitsInfo, commit{
			Message:       commitInf.Message,
			Oid:           commitInf.Oid,
			CommittedDate: commitInf.CommittedDate,
			Author:        commitInf.Author.User.Login,
		})
	}

	result := &CommitsResponseContract{
		Commits: commitsInfo,
	}

	return result, nil
}

func gitlabGetCommits(name string, owner string, accessToken string) (*CommitsResponseContract, error) {

	projectName := owner + "/" + name
	project, _, err := gitlabClient.Projects.GetProject(projectName, nil)
	if err != nil {
		return nil, err
	}

	all := true
	opts := gitlab.ListCommitsOptions{
		All: &all,
	}

	commitsData, _, err := gitlabClient.Commits.ListCommits(project.ID, &opts)
	if err != nil {
		return nil, err
	}

	strFormatDate := env.Get().DefaultConfiguration.DateFormat
	commits := []commit{}
	for _, c := range commitsData {
		theData := commit{}

		theData.Message = c.Message
		theData.Author = c.AuthorEmail
		theData.Oid = c.ID

		if c.CommittedDate != nil {
			theData.CommittedDate = c.CommittedDate.Format(strFormatDate)
		}

		commits = append(commits, theData)
	}

	result := &CommitsResponseContract{
		Commits: commits,
	}

	return result, nil
}
