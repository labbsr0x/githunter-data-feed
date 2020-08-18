package services

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	gitlabService "github.com/labbsr0x/githunter-api/services/gitlab"
	"github.com/xanzy/go-gitlab"
)

// ContractInterface
type Contract interface {
	GetLastRepos(string, string) (*ReposResponseContract, error)
	GetInfoCodePage(string, string, string, string) (*CodeResponseContract, error)
	GetCommitsRepo(string, string, string, string) (*CommitsResponseContract, error)
	GetIssues(string, string, string, string) (*IssuesResponseContract, error)
	GetPulls(string, string, string, string) (*PullsResponseContract, error)
}

type defaultContract struct{}

type participants struct {
	TotalCount int      `json:"totalCount"`
	User       []string `json:"users"`
}

type comments struct {
	TotalCount int       `json:"totalCount"`
	UpdatedAt  string    `json:"updatedAt"`
	Data       []comment `json:"data"`
}

type comment struct {
	CreatedAt string `json:"createdAt"`
	Author    string `json:"author"`
}

func New() Contract {
	return &defaultContract{}
}

func (d *defaultContract) Gitlab(accessToken string) *gitlabService.Gitlab {
	return gitlabService.New(accessToken)
}

var gitlabClient *gitlab.Client

func gitlabNewClient(accessToken string) *gitlab.Client {
	client, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(env.Get().ApiGitlabURL))
	if err != nil {
		return nil
	}
	return client
}
