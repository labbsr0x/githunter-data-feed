package services

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
	GetInfoCodePage(string, string, string, string) (*CodeResponseContract, error)
}

type defaultContract struct{}

func New() Contract {
	return &defaultContract{}
}
