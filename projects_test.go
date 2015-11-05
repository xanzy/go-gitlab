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
		t.Fatalf("Expected error nil, got: %v", err)
	}

	count := len(projects)
	if count != 2 {
		t.Errorf("Expected number of projects %q, got %q", 2, count)
	}
	if projects[0].Name != "project" {
		t.Errorf("Expected project name %q, got %q", "project", projects[0].Name)
	}
	if projects[1].Name != "project2" {
		t.Errorf("Expected project name %q, got %q", "project2", projects[1].Name)
	}
}

func TestGetProject(t *testing.T) {
	ts, client := Stub("stubs/projects/show.json")
	defer ts.Close()

	project, _, err := client.Projects.GetProject(1)

	if err != nil {
		t.Fatalf("Expected error nil, got: %v", err)
	}

	if project.Name != "project" {
		t.Errorf("Expected project name %q, got %q", "project", project.Name)
	}
	if project.Namespace.Name != "group" {
		t.Errorf("Expected namespace name %q, got %q", "group", project.Namespace.Name)
	}
	if project.Permissions.ProjectAccess.AccessLevel != MasterPermissions {
		t.Errorf("Expected project access level %q, got %q", MasterPermissions, project.Permissions.ProjectAccess.AccessLevel)
	}
	if project.Permissions.GroupAccess != nil {
		t.Errorf("Expected project group access nil, got %q", project.Permissions.GroupAccess)
	}
}
