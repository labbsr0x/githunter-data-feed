package services

import (
	"fmt"
	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

type OrganizationResponseContract struct {
	Members membersStruct `json:"members"`
}

type membersStruct struct {
	Total   int      `json:"total"`
	Members []string `json:"data"`
}

func (d *defaultContract) GetMembers(organization string, provider string, accessToken string) (*OrganizationResponseContract, error) {
	theContract := &OrganizationResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api")

	switch provider {
	case `github`:
		theContract, err = githubGetMembers(organization, accessToken)
		break
	case `gitlab`:

		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetMembers unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetMembers returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil

}

func githubGetMembers(organization string, accessToken string) (*OrganizationResponseContract, error) {

	members := []string{}
	after := ""
	for {
		githubMembers, err := github.GetMembers(organization, after, accessToken)
		if err != nil {
			return nil, err
		}

		data := []string{}
		data, after = formatContractMembers4Github(githubMembers)

		members = append(members, data...)

		if after == "" {
			break
		}
	}

	result := &OrganizationResponseContract{
		Members: membersStruct{
			Total:   len(members),
			Members: members,
		},
	}

	return result, nil
}

func formatContractMembers4Github(response *github.Response) ([]string, string) {

	lastCursor := ""
	data := []string{}
	for _, v := range response.Organization.MembersWithRole.Edges {
		lastCursor = v.Cursor
		data = append(data, v.Node.Login)
	}

	return data, lastCursor
}
