package services

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
	GetCommitsRepo(string, string, int, string, string) (*CommitsResponseContract, error)
}

type defaultContract struct{}

func New() Contract {
	return &defaultContract{}
}
