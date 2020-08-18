package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

func (g *Gitlab) ListProjectMergeRequests(state string, projectID int) ([]*gitlab.MergeRequest, error) {
	opts := gitlab.ListProjectMergeRequestsOptions{
		State: &state,
	}

	mergeRequests, _, err := g.Client.MergeRequests.ListProjectMergeRequests(projectID, &opts)
	if err != nil {
		return nil, err
	}

	return mergeRequests, nil
}

func (g *Gitlab) GetMergeRequestParticipants(projectID int, mergeRequestID int) ([]*gitlab.BasicUser, error) {

	participants, _, err := g.Client.MergeRequests.GetMergeRequestParticipants(projectID, mergeRequestID)
	if err != nil {
		return nil, err
	}

	return participants, nil
}

func (g *Gitlab) GetProjectDescription(projectName string) (*gitlab.Project, error) {
	project, _, err := g.Client.Projects.GetProject(projectName, nil)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (g *Gitlab) GetDiscussions(projectID int, mergeRequestID int) ([]*gitlab.Discussion, error) {
	discussions, _, err := g.Client.Discussions.ListMergeRequestDiscussions(projectID, mergeRequestID, nil)
	if err != nil {
		return nil, err
	}

	return discussions, nil
}
