package services

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
	GetIssues(int, string, string, string, string) (*IssuesResponseContract, error)
}

type defaultContract struct{}

func New() Contract {
	return &defaultContract{}
}
