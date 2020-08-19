package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-api/infra/env"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/labbsr0x/githunter-api/services/gitlab"
	"github.com/sirupsen/logrus"
	gitlabLib "github.com/xanzy/go-gitlab"
)

type PullsResponseContract struct {
	Total int    `json:"total"`
	Pulls []pull `json:"data"`
}

type pull struct {
	Number       int          `json:"number"`
	State        string       `json:"state"`
	CreatedAt    string       `json:"createdAt"`
	ClosedAt     string       `json:"closedAt"`
	Merged       bool         `json:"merged"`
	MergedAt     string       `json:"mergedAt"`
	Author       string       `json:"author"`
	Labels       []string     `json:"labels"`
	Comments     comments     `json:"comments"`
	Participants participants `json:"participants"`
}

func (d *defaultContract) GetPulls(owner string, name string, provider string, accessToken string) (*PullsResponseContract, error) {
	theContract := &PullsResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetPulls(owner, name, accessToken)
		break
	case `gitlab`:
		client = d.Gitlab(accessToken)
		theContract, err = gitlabGetPulls(owner, name)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetPulls unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetPulls returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetPulls(owner string, repo string, accessToken string) (*PullsResponseContract, error) {
	pullsOpened, err := github.GetPulls(owner, repo, accessToken, false)

	if err != nil {
		return nil, err
	}

	data := formatContract4Github(pullsOpened)

	pullsClosed, err := github.GetPulls(owner, repo, accessToken, true)

	if err != nil {
		return nil, err
	}

	data = append(data, formatContract4Github(pullsClosed)...)

	result := &PullsResponseContract{
		Total: pullsOpened.Repository.Pulls.TotalCount + pullsClosed.Repository.Pulls.TotalCount,
		Pulls: data,
	}

	return result, nil
}

func formatContract4Github(response *github.Response) []pull {

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

//Gitlab Session
var client *gitlab.Gitlab

func gitlabGetPulls(owner string, name string) (*PullsResponseContract, error) {

	projectName := owner + "/" + name
	project, err := client.GetProjectDescription(projectName)
	if err != nil {
		return nil, err
	}

	mergeRequests, _ := client.ListProjectMergeRequests("merged", project.ID)

	total := len(mergeRequests)

	data := formatContract4Gitlab(mergeRequests)

	mergeRequests, err = client.ListProjectMergeRequests("closed", project.ID)
	if err != nil {
		return nil, err
	}
	total = total + len(mergeRequests)

	data = append(data, formatContract4Gitlab(mergeRequests)...)

	mergeRequests, err = client.ListProjectMergeRequests("opened", project.ID)
	if err != nil {
		return nil, err
	}
	total = total + len(mergeRequests)

	data = append(data, formatContract4Gitlab(mergeRequests)...)

	data = fillDiscussion(data, project.ID)
	data = fillParticipants(data, project.ID)

	result := &PullsResponseContract{
		Total: total,
		Pulls: data,
	}

	return result, nil
}

func formatContract4Gitlab(mergeRequests []*gitlabLib.MergeRequest) []pull {

	strFormatDate := env.Get().DefaultConfiguration.DateFormat

	data := []pull{}
	for _, v := range mergeRequests {
		theData := pull{}
		theData.Number = v.IID
		theData.State = v.State
		theData.Author = v.Author.Username

		if v.CreatedAt != nil {
			theData.CreatedAt = v.CreatedAt.Format(strFormatDate)
		}

		if v.ClosedAt != nil {
			theData.ClosedAt = v.ClosedAt.Format(strFormatDate)
		}
		if v.MergedAt != nil {
			theData.MergedAt = v.MergedAt.Format(strFormatDate)
		}

		theData.Merged = false
		if v.State == "merged" {
			theData.Merged = true
		}

		for _, l := range v.Labels {
			theData.Labels = append(theData.Labels, l)
		}

		data = append(data, theData)
	}
	return data
}

func fillDiscussion(mergeRequests []pull, projectID int) []pull {

	strFormatDate := env.Get().DefaultConfiguration.DateFormat

	mergeRequestsWithDiscussion := []pull{}
	for _, v := range mergeRequests {

		discussions := []comment{}

		lastUpdated := ""

		mrDiscussions, err := client.GetDiscussions(projectID, v.Number)
		if err != nil {
			continue
		}
		for _, d := range mrDiscussions {

			if len(d.Notes) > 0 && d.Notes[0].UpdatedAt != nil {
				lastUpdated = d.Notes[0].UpdatedAt.Format(strFormatDate)
			}

			closedAt := ""
			if len(d.Notes) > 0 && d.Notes[0].CreatedAt != nil {
				closedAt = d.Notes[0].CreatedAt.Format(strFormatDate)
			}

			author := ""
			if len(d.Notes) > 0 && d.Notes[0].Author.Username != "" {
				author = d.Notes[0].Author.Username
			}
			theComment := comment{
				CreatedAt: closedAt,
				Author:    author,
			}

			discussions = append(discussions, theComment)
		}

		v.Comments = comments{
			TotalCount: len(discussions),
			UpdatedAt:  lastUpdated,
			Data:       discussions,
		}
		mergeRequestsWithDiscussion = append(mergeRequestsWithDiscussion, v)

	}

	return mergeRequestsWithDiscussion
}

func fillParticipants(mergeRequests []pull, projectID int) []pull {

	mergeRequestsWithParticipants := []pull{}

	for _, v := range mergeRequests {

		listOfUsers := []string{}
		mrParticipants, err := client.GetMergeRequestParticipants(projectID, v.Number)
		if err != nil {
			continue
		}
		for _, p := range mrParticipants {
			listOfUsers = append(listOfUsers, p.Username)
		}

		v.Participants = participants{
			TotalCount: len(listOfUsers),
			User:       listOfUsers,
		}
		mergeRequestsWithParticipants = append(mergeRequestsWithParticipants, v)

	}

	return mergeRequestsWithParticipants
}
