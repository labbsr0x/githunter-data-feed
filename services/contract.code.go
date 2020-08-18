package services

import (
	"fmt"
	"math"

	"github.com/labbsr0x/githunter-api/services/github"
	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

//struct response of Code page
type CodeResponseContract struct {
	Name                 string     `json:"name"`
	Description          string     `json:"description"`
	CreatedAt            string     `json:"createdAt"`
	PrimaryLanguage      string     `json:"primaryLanguage"`
	RepositoryTopics     []string   `json:"repositoryTopics"`
	Watchers             int        `json:"watchers"`
	Stars                int        `json:"stars"`
	Forks                int        `json:"forks"`
	LastCommitDate       string     `json:"lastCommitDate"`
	Commits              int        `json:"commits"`
	HasHomepageUrl       bool       `json:"hasHomepageUrl"`
	HasReadmeFile        bool       `json:"hasReadmeFile"`
	HasContributingFile  bool       `json:"hasContributingFile"`
	LicenseInfo          string     `json:"licenseInfo"`
	HasCodeOfConductFile bool       `json:"hasCodeOfConductFile"`
	Releases             int        `json:"releases"`
	Contributors         int        `json:"contributors"` //*Info no longer available on Github*
	Languages            *languages `json:"languages"`
	DiskUsage            int        `json:"diskUsage"`
}

type codeOfConduct struct {
	ResourcePath string `json:"resourcePath"`
}

type languages struct {
	Quantity  int        `json:"quantity"`
	Languages []language `json:"languages"`
}

type language struct {
	Size float64 `json:"size"`
	Name string  `json:"name"`
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
		client = d.Gitlab(accessToken)
		theContract, err = gitlabGetCodePageInfo(nameRepo, ownerRepo)
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

	codeOfConduct := codeOfConduct{
		code.Viewer.CodeOfConduct.ResourcePath,
	}

	langsInfo := []language{}
	for _, lang := range code.Viewer.Languages.Languages {
		langsInfo = append(langsInfo, language{
			Size: float64(lang.Size),
			Name: lang.Language.Name,
		})
	}

	languages := &languages{
		Quantity:  code.Viewer.Languages.Quantity,
		Languages: langsInfo,
	}

	hasHomepageUrl := false
	if len(code.Viewer.HomepageUrl) > 0 {
		hasHomepageUrl = true
	}

	hasReadmeFile := false
	if code.Viewer.Readme.ByteSize > 0 {
		hasReadmeFile = true
	}

	hasContributingFile := false
	if code.Viewer.Contributing.ByteSize > 0 {
		hasContributingFile = true
	}

	hasCodeOfConductFile := false
	if len(codeOfConduct.ResourcePath) > 0 {
		hasCodeOfConductFile = true
	}

	result := &CodeResponseContract{
		Name:                 code.Viewer.Name,
		Description:          code.Viewer.Description,
		CreatedAt:            code.Viewer.CreatedAt,
		PrimaryLanguage:      code.Viewer.PrimaryLanguage.Name,
		RepositoryTopics:     topics,
		Watchers:             code.Viewer.Watchers.TotalCount,
		Stars:                code.Viewer.Stars.TotalCount,
		Forks:                code.Viewer.Forks,
		LastCommitDate:       code.Viewer.LastCommit.DefaultBranch.LastCommitDate,
		Commits:              code.Viewer.LastCommit.DefaultBranch.CommitsQuanity.TotalCount,
		HasHomepageUrl:       hasHomepageUrl,
		HasReadmeFile:        hasReadmeFile,
		HasContributingFile:  hasContributingFile,
		LicenseInfo:          code.Viewer.LicenseInfo.Name,
		HasCodeOfConductFile: hasCodeOfConductFile,
		Releases:             code.Viewer.Releases.TotalCount,
		// Contributors:         code.Viewer.Contributors.History.TotalCount, *Info no longer available*
		Languages: languages,
		DiskUsage: code.Viewer.DiskUsage,
	}

	return result, nil
}

// Gitlab REST
func gitlabGetCodePageInfo(nameRepo string, ownerRepo string) (*CodeResponseContract, error) {
	projectPath := ownerRepo + "/" + nameRepo
	project, err := client.GetProjectInfo(projectPath)
	if err != nil {
		return nil, err
	}

	name := project.Name
	description := project.Description
	createAt := project.CreatedAt
	repositoryTopics := []string{""}
	watchers := 0
	stars := project.StarCount
	forks := project.ForksCount
	lastCommitDate := project.LastActivityAt

	commitsQuantity := 0
	_, resp, err := client.Client.Commits.ListCommits(projectPath, nil)
	if resp.Response.StatusCode == 200 {
		commitsQuantity = resp.TotalItems
	}

	// commitsQuantity := project.Statistics.CommitCount

	hasHomepageUrl := false

	hasReadmeFile := false
	if project.ReadmeURL != "" {
		hasReadmeFile = true
	}

	defaultBranch := project.DefaultBranch
	opts := gitlab.GetFileOptions{
		Ref: &defaultBranch,
	}

	hasContributingFile := false
	contributingFile, resp, err := client.Client.RepositoryFiles.GetFile(projectPath, "CONTRIBUTING.md", &opts)
	if contributingFile != nil {
		hasContributingFile = true
	}

	getPopularLicense := true
	licensesOpts := gitlab.ListLicenseTemplatesOptions{
		Popular: &getPopularLicense,
	}
	licenses, resp, err := client.Client.LicenseTemplates.ListLicenseTemplates(&licensesOpts, nil)
	licenseIndex := licenses[0]
	license := licenseIndex.Name

	hasCodeOfConductFile := false
	codeOfConductFile, resp, err := client.Client.RepositoryFiles.GetFile(projectPath, "CODE_OF_CONDUCT.md", &opts)
	if codeOfConductFile != nil {
		hasCodeOfConductFile = true
	}

	releasesQuantity := 0
	_, resp, err = client.Client.Releases.ListReleases(projectPath, nil)
	if resp.Response.StatusCode == 200 {
		releasesQuantity = resp.TotalItems
	}

	contributors := 0
	_, resp, err = client.Client.ProjectMembers.ListAllProjectMembers(projectPath, nil)
	if resp.Response.StatusCode == 200 {
		contributors = resp.TotalItems
	}

	languagesResp, resp, err := client.Client.Projects.GetProjectLanguages(projectPath, nil)

	langsInfo := []language{}
	primaryLanguage := ""
	if resp.Response.StatusCode == 200 {
		index := 0
		for key, value := range *languagesResp {
			if index == 0 {
				primaryLanguage = key
			}
			vf64 := float64(value)
			langsInfo = append(langsInfo, language{
				Size: (math.Floor(vf64*100) / 100),
				Name: key,
			})
			index++
		}
	}

	langs := &languages{
		Quantity:  len(*languagesResp),
		Languages: langsInfo,
	}

	strFormatDate := "2006-01-02T15:04:05Z"
	result := &CodeResponseContract{
		Name:                 name,
		Description:          description,
		CreatedAt:            createAt.Format(strFormatDate),
		PrimaryLanguage:      primaryLanguage,
		RepositoryTopics:     repositoryTopics,
		Watchers:             watchers,
		Stars:                stars,
		Forks:                forks,
		LastCommitDate:       lastCommitDate.Format(strFormatDate),
		Commits:              commitsQuantity,
		HasHomepageUrl:       hasHomepageUrl,
		HasReadmeFile:        hasReadmeFile,
		HasContributingFile:  hasContributingFile,
		LicenseInfo:          license,
		HasCodeOfConductFile: hasCodeOfConductFile,
		Releases:             releasesQuantity,
		Contributors:         contributors,
		Languages:            langs,
		DiskUsage:            0,
	}

	return result, nil
}
