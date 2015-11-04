package gitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAllProjects(t *testing.T) {
	ts, client := Stub("stubs/projects/index.json")

	opt := &ListProjectsOptions{}
	projects, _, err := client.Projects.ListAllProjects(opt)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(projects), 2)
	assert.Equal(t, projects[0].Name, "project")
	assert.Equal(t, projects[1].Name, "project2")
	defer ts.Close()
}

func TestGetProject(t *testing.T) {
	ts, client := Stub("stubs/projects/show.json")

	project, _, err := client.Projects.GetProject(1)

	assert.Equal(t, err, nil)
	assert.Equal(t, project.Name, "project")
	assert.Equal(t, project.Namespace.Name, "group")
	assert.Equal(t, project.Permissions.ProjectAccess.AccessLevel, MasterPermissions)
	assert.Nil(t, project.Permissions.GroupAccess)
	defer ts.Close()
}