package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

type IssuesResponseContract struct {
	Issues []issue `json:"issues"`
}

type issue struct {
	Number       int          `json:"number"`
	State        string       `json:"state"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
	ClosedAt     string       `json:"closedAt"`
	Author       string       `json:"author"`
	Labels       []string     `json:"labels"`
	Comments     comments     `json:"comments"`
	Participants participants `json:"participants"`
}

func (d *defaultContract) GetIssues(owner string, repo string, provider string, accessToken string) (*IssuesResponseContract, error) {
	theContract := &IssuesResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetIssues(owner, repo, accessToken)
		break
	case `gitlab`:
		gitlabClient = gitlabNewClient(accessToken)
		theContract, err = gitlabGetIssues(owner, repo, accessToken)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetIssues unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetIssues returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetIssues(owner string, repo string, accessToken string) (*IssuesResponseContract, error) {
	issuesResp, err := github.GetIssues(owner, repo, accessToken)

	if err != nil {
		return nil, err
	}

	issues := formatContractIssuesGithub(issuesResp)

	result := &IssuesResponseContract{
		Issues: issues,
	}

	return result, nil
}

func formatContractIssuesGithub(issuesResp *github.Response) []issue {

	issues := []issue{}
	for _, v := range issuesResp.Repository.Issues.Nodes {
		theIssue := issue{}
		theIssue.Number = v.Number
		theIssue.State = v.State
		theIssue.Author = v.Author.Login
		theIssue.CreatedAt = v.CreatedAt
		theIssue.UpdatedAt = v.UpdatedAt
		theIssue.ClosedAt = v.ClosedAt

		theIssue.Participants.TotalCount = v.Participants.TotalCount
		for _, t := range v.Participants.User {
			theIssue.Participants.User = append(theIssue.Participants.User, t.Login)
		}

		for _, l := range v.Labels.Label {
			theIssue.Labels = append(theIssue.Labels, l.Name)
		}

		theIssue.Comments.TotalCount = v.TimelineItems.TotalCount

		for _, t := range v.TimelineItems.Data {
			theComment := comment{}
			theComment.Author = t.Author.Login
			theComment.CreatedAt = t.CreatedAt
			theIssue.Comments.Data = append(theIssue.Comments.Data, theComment)
		}

		issues = append(issues, theIssue)
	}

	return issues
}

func gitlabGetIssues(owner string, repo string, accessToken string) (*IssuesResponseContract, error) {

	projectName := owner + "/" + repo
	project, _, err := gitlabClient.Projects.GetProject(projectName, nil)
	if err != nil {
		return nil, err
	}

	issuesData, _, err := gitlabClient.Issues.ListProjectIssues(project.ID, nil)
	if err != nil {
		return nil, err
	}

	issues := []issue{}

	for _, i := range issuesData {
		theIssue := issue{}

		theIssue.Number = i.IID
		theIssue.State = i.State
		theIssue.Author = i.Author.Username

		if i.CreatedAt != nil {
			theIssue.CreatedAt = i.CreatedAt.String()
		}

		if i.UpdatedAt != nil {
			theIssue.UpdatedAt = i.UpdatedAt.String()
		}

		if i.ClosedAt != nil {
			theIssue.ClosedAt = i.ClosedAt.String()
		}

		//TODO implements get participants https://github.com/xanzy/go-gitlab/pull/920

		for _, l := range i.Labels {
			theIssue.Labels = append(theIssue.Labels, l)
		}

		notes, resp, err := gitlabClient.Notes.ListIssueNotes(project.ID, i.IID, nil)
		if err != nil {
			return nil, err
		}

		theIssue.Comments.TotalCount = resp.TotalItems

		for _, n := range notes {
			theComment := comment{}
			theComment.Author = n.Author.Username

			if n.CreatedAt != nil {
				theComment.CreatedAt = n.CreatedAt.String()
			}

			theIssue.Comments.Data = append(theIssue.Comments.Data, theComment)
		}

		issues = append(issues, theIssue)
	}

	result := &IssuesResponseContract{
		Issues: issues,
	}

	return result, nil
}
