package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

type CommentsResponseContract struct {
	Data []comment `json:"data"`
}

type comment struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"createdAt"`
	Number    int    `json:"number"`
	Url       string `json:"url"`
	Id        string `json:"id"`
	Author    string `json:"author"`
}

func (d *defaultContract) GetComments(ids []string, provider string, accessToken string) (*CommentsResponseContract, error) {
	theContract := &CommentsResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetComments(ids, accessToken)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetComments unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetComments returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetComments(ids []string, accessToken string) (*CommentsResponseContract, error) {
	commentsResp, err := github.GetComments(ids, accessToken)

	if err != nil {
		return nil, err
	}

	commentsData := formatContractCommentsGithub(commentsResp)

	result := &CommentsResponseContract{
		Data: commentsData,
	}

	return result, nil
}

func formatContractCommentsGithub(commentsResp *github.CommentsResponse) []comment {

	comments := []comment{}
	for _, v := range commentsResp.Nodes {
		theComment := comment{}
		theComment.Number = v.Issue.Number
		theComment.Author = v.Author.Login
		theComment.CreatedAt = v.CreatedAt
		theComment.Name = v.Repository.Name
		theComment.Owner = v.Repository.Owner.Login
		theComment.Url = v.Url
		theComment.Id = v.Id

		comments = append(comments, theComment)
	}

	return comments
}
