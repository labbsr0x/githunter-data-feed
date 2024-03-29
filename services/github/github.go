package github

type Response struct {
	Repository   repository   `json:"repository"`
	Organization organization `json:"organization"`
}

type repository struct {
	Issues issue `json:"issues"`
	Pulls  pulls `json:"pullRequests"`
}

type participants struct {
	TotalCount int    `json:"totalCount"`
	User       []user `json:"nodes"`
}

type user struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type author struct {
	User user `json:"user"`
}

type labels struct {
	Label []label `json:"nodes"`
}

type label struct {
	Name string `json:"name"`
}

type comments struct {
	TotalCount int            `json:"totalCount"`
	Data       []shortComment `json:"nodes"`
}

type shortComment struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	CreatedAt string `json:"createdAt"`
	Author    user   `json:"author"`
}

type node struct {
	Name string `json:"name"`
}

type count struct {
	TotalCount int `json:"totalCount"`
}

type defaultBranch struct {
	DefaultBranch nodeTarget `json:"target"`
}

type repositoryTopics struct {
	Nodes []topic `json:"nodes"`
}

type topic struct {
	Name node `json:"topic"`
}

type text struct {
	Text string `json:"text"`
}

type byteSize struct {
	ByteSize int `json:"byteSize"`
}

type language struct {
	Name string `json:"language"`
}

type pathRepo struct {
	Name  string `json:"name"`
	Owner owner  `json:"owner"`
}

type owner struct {
	Login string `json:"login"`
}

type followers struct {
	Follower []loginUser `json:"nodes"`
}

type loginUser struct {
	Login string `json:"login"`
}
