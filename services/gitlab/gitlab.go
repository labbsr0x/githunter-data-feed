package gitlab

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	gitlab "github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	Client *gitlab.Client
}

func New(accessToken string) *Gitlab {
	return &Gitlab{
		Client: client(accessToken),
	}
}

func (g *Gitlab) GetProjectInfo(projectName string) (*gitlab.Project, error) {
	project, _, err := g.Client.Projects.GetProject(projectName, nil)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func client(accessToken string) *gitlab.Client {
	client, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(env.Get().GitlabGraphQLURL))
	if err != nil {
		return nil
	}
	return client
}
