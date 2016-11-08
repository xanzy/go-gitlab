package gitlab

import (
	"fmt"
	"net/url"
)

// VariablesService handles communication with the project variables related methods
// of the Gitlab API
//
// Gitlab API Docs : https://docs.gitlab.com/ce/api/build_variables.html
type VariablesService struct {
	client *Client
}

//Variable represents a variable available for each build of the given project
//
// Gitlab API Docs : https://docs.gitlab.com/ce/api/build_variables.html
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ListVariables gets the a list of project variables in a project
//
// Gitlab API Docs:
// https://docs.gitlab.com/ce/api/build_variables.html#list-project-variables
func (s *VariablesService) ListVariables(pid interface{}) ([]*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/variables", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var variables []*Variable
	resp, err := s.client.Do(req, &variables)
	if err != nil {
		return nil, resp, err
	}

	return variables, resp, err
}

// GetSingleVariable gets a single project variable of a project
//
// Gitlab API Docs:
// https://docs.gitlab.com/ce/api/build_variables.html#show-variable-details
func (s *VariablesService) GetSingleVariable(pid interface{}, variableKey string) (*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/variables/%s", url.QueryEscape(project), variableKey)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	variable := new(Variable)
	resp, err := s.client.Do(req, variable)
	if err != nil {
		return nil, resp, err
	}

	return variable, resp, err
}

// CreateVariable creates a variable for a given project
//
// Gitlab API Docs:
// https://docs.gitlab.com/ce/api/build_variables.html#create-variable
func (s *VariablesService) CreateVariable(pid interface{}, variable Variable) (*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/variables", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, variable)
	if err != nil {
		return nil, nil, err
	}

	createdVariable := new(Variable)
	resp, err := s.client.Do(req, createdVariable)
	if err != nil {
		return nil, resp, err
	}

	return createdVariable, resp, err
}

// UpdateVariable updates an existing project variable
// The variable key must exist
//
// Gitlab API Docs:
// https://docs.gitlab.com/ce/api/build_variables.html#update-variable
func (s *VariablesService) UpdateVariable(pid interface{}, variableKey string, variable Variable) (*Variable, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/variables/%s", url.QueryEscape(project), variableKey)

	req, err := s.client.NewRequest("PUT", u, variable)
	if err != nil {
		return nil, nil, err
	}

	updatedVariable := new(Variable)
	resp, err := s.client.Do(req, updatedVariable)
	if err != nil {
		return nil, resp, err
	}

	return updatedVariable, resp, err
}

// RemoveVariable removes a project variable of a given project identified by its key
//
// Gitlab API Docs:
// https://docs.gitlab.com/ce/api/build_variables.html#remove-variable
func (s *VariablesService) RemoveVariable(pid interface{}, variableKey string) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("projects/%s/variables/%s", url.QueryEscape(project), variableKey)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
