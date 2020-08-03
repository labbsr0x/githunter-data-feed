package services

// ContractInterface
type Contract interface {
	GetLastRepos(int, string, string) (*ReposResponseContract, error)
}

type defaultContract struct{}

func New() Contract {
	return &defaultContract{}
}
