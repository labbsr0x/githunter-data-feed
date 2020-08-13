package services

import "github.com/labbsr0x/githunter-api/services/gitlab"

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
	GetInfoCodePage(string, string, string, string) (*CodeResponseContract, error)
	GetCommitsRepo(int, string, string, string, string) (*CommitsResponseContract, error)
	GetIssues(int, string, string, string, string) (*IssuesResponseContract, error)
	GetPulls(int, string, string, string, string) (*PullsResponseContract, error)
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

func (d *defaultContract) Gitlab(accessToken string) *gitlab.Gitlab {
	return gitlab.New(accessToken)
}
