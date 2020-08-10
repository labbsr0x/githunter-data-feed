package services

import (
	"fmt"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
)

//struct response of Code page
type CodeResponseContract struct {
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	CreatedAt        string        `json:"createdAt"`
	PrimaryLanguage  string        `json:"primaryLanguage"`
	RepositoryTopics []string      `json:"repositoryTopics"`
	Watchers         int           `json:"watchers"`
	Stars            int           `json:"stars"`
	Forks            int           `json:"forks"`
	LastCommitDate   string        `json:"lastCommitDate"`
	Commits          int           `json:"commits"`
	Readme           string        `json:"readme"`
	Contributing     string        `json:"contributing"`
	LicenseInfo      string        `json:"licenseInfo"`
	CodeOfConduct    CodeOfConduct `json:"codeOfConduct"`
	Releases         int           `json:"releases"`
	Contributors     int           `json:"contributors"`
	Languages        *Languages    `json:"languages"`
	DiskUsage        int           `json:"diskUsage"`
}

type CodeOfConduct struct {
	Body         string `json:"body"`
	ResourcePath string `json:"resourcePath"`
}

type Languages struct {
	Quantity  int        `json:"quantity"`
	Languages []Language `json:"languages"`
}

type Language struct {
	Size int    `json:"size"`
	Name string `json:"name"`
}

func (d *defaultContract) GetInfoCodePage(nameRepo string, ownerRepo string, accessToken string, provider string) (*CodeResponseContract, error) {

	theContract := &CodeResponseContract{}
	var err error

	logrus.WithFields(logrus.Fields{
		"provider": provider,
	}).Debug("Making a request in an external api for get the last repositories of the user")

	switch provider {
	case `github`:
		theContract, err = githubGetCodePageInfo(nameRepo, ownerRepo, accessToken)
		break
	case `gitlab`:
		// theContract, err = gitlabGetLastRepos(numberOfRepos, accessToken)
		break
	case ``:
		//TODO: Call all providers
		break
	default:
		return nil, fmt.Errorf("GetCodePageInfo unknown provider: %s", provider)
	}

	if theContract == nil {
		logrus.Debug("GetCodePageInfo returned a null answer")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"provider": provider,
		}).Warn(err.Error())
	}

	return theContract, nil
}

func githubGetCodePageInfo(nameRepo string, ownerRepo string, accessToken string) (*CodeResponseContract, error) {
	code, err := github.GetInfoCodePage(nameRepo, ownerRepo, accessToken)
	if err != nil {
		return nil, err
	}

	topics := []string{}
	for _, topic := range code.Viewer.RepositoryTopics.Nodes {
		topics = append(topics, topic.Name.Name)
	}

	codeOfConduct := CodeOfConduct{
		code.Viewer.CodeOfConduct.Body,
		code.Viewer.CodeOfConduct.ResourcePath,
	}

	langsInfo := []Language{}
	for _, lang := range code.Viewer.Languages.Languages {
		langsInfo = append(langsInfo, Language{
			Size: lang.Size,
			Name: lang.Language.Name,
		})
	}

	languages := &Languages{
		Quantity:  code.Viewer.Languages.Quantity,
		Languages: langsInfo,
	}

	result := &CodeResponseContract{
		Name:             code.Viewer.Name,
		Description:      code.Viewer.Description,
		CreatedAt:        code.Viewer.CreatedAt,
		PrimaryLanguage:  code.Viewer.PrimaryLanguage.Name,
		RepositoryTopics: topics,
		Watchers:         code.Viewer.Watchers.TotalCount,
		Stars:            code.Viewer.Stars.TotalCount,
		Forks:            code.Viewer.Forks,
		LastCommitDate:   code.Viewer.LastCommit.DefaultBranch.LastCommitDate,
		Commits:          code.Viewer.LastCommit.DefaultBranch.CommitsQuanity.TotalCount,
		Readme:           code.Viewer.Readme.Text,
		Contributing:     code.Viewer.Contributing.Text,
		LicenseInfo:      code.Viewer.LicenseInfo.Name,
		CodeOfConduct:    codeOfConduct,
		Releases:         code.Viewer.Releases.TotalCount,
		Contributors:     code.Viewer.Contributors.History.TotalCount,
		Languages:        languages,
		DiskUsage:        code.Viewer.DiskUsage,
	}

	return result, nil
}

func gitlabGetPageCodePageInfo(nameRepo string, ownerRepo string, accessToken string) (*CodeResponseContract, error) {
	// code, err := gitlab.GetInfoCodePage(nameRepo, ownerRepo, accessToken)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
