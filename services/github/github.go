package github

type IssuesResponse struct {
	Repository repository `json:"repository"`
}

type repository struct {
	Issues issue `json:"issues"`
	Pulls pulls `json:"pullRequests"`
}

type participants struct {
	TotalCount int `json:"totalCount"`
	User []user `json:"nodes"`
}

type user struct {
	Login string `json:"login"`
	Name string `json:"name"`
}

type labels struct {
	Label []label `json:"nodes"`
}

type label struct {
	Name string `json:"name"`
}

type comments struct {
	TotalCount int       `json:"totalCount"`
	Data      []comment `json:"nodes"`
}

type comment struct {
	CreatedAt string 	`json:"createdAt"`
	Author    user 		`json:"author"`
}