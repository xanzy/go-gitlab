package gitlab

import (
	"testing"
)

func TestListAllProjects(t *testing.T) {
	ts, client := Stub("stubs/projects/index.json")
	defer ts.Close()

	opt := &ListProjectsOptions{}
	projects, _, err := client.Projects.ListAllProjects(opt)

	if err != nil {
		t.Fatal(err)
	}

	if len(projects) != 2 {
		t.Fail()
	}
	if projects[0].Name != "project" {
		t.Fail()
	}
	if projects[1].Name != "project2" {
		t.Fail()
	}
}

func TestGetProject(t *testing.T) {
	ts, client := Stub("stubs/projects/show.json")
	defer ts.Close()

	project, _, err := client.Projects.GetProject(1)

	if err != nil {
		t.Fatal(err)
	}

	if project.Name != "project" {
		t.Fail()
	}
	if project.Namespace.Name != "group" {
		t.Fail()
	}
	if project.Permissions.ProjectAccess.AccessLevel != MasterPermissions {
		t.Fail()
	}
	if project.Permissions.GroupAccess != nil {
		t.Fail()
	}
}
