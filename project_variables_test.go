package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectVariablesService_ListVariables(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"key": "TEST_VARIABLE_1",
					"variable_type": "env_var",
					"value": "TEST_1"
				}
			]
		`)
	})

	want := []*ProjectVariable{{
		Key:              "TEST_VARIABLE_1",
		Value:            "TEST_1",
		VariableType:     "env_var",
		Protected:        false,
		Masked:           false,
		EnvironmentScope: "",
	}}

	pvs, resp, err := client.ProjectVariables.ListVariables(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pvs)

	pvs, resp, err = client.ProjectVariables.ListVariables(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pvs)

	pvs, resp, err = client.ProjectVariables.ListVariables(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pvs)

	pvs, resp, err = client.ProjectVariables.ListVariables(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pvs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectVariablesService_GetVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/variables/TEST_VARIABLE_1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "filter%5Benvironment_scope%5D=prod")
		fmt.Fprintf(w, `
			{
				"key": "TEST_VARIABLE_1",
				"variable_type": "env_var",
				"value": "TEST_1",
				"protected": false,
				"masked": true
			}
		`)
	})

	want := &ProjectVariable{
		Key:              "TEST_VARIABLE_1",
		Value:            "TEST_1",
		VariableType:     "env_var",
		Protected:        false,
		Masked:           true,
		EnvironmentScope: "",
	}

	pv, resp, err := client.ProjectVariables.GetVariable(1, "TEST_VARIABLE_1", &GetProjectVariableOptions{Filter: &VariableFilter{EnvironmentScope: "prod"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pv)

	pv, resp, err = client.ProjectVariables.GetVariable(1.01, "TEST_VARIABLE_1", nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.GetVariable(1, "TEST_VARIABLE_1", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.GetVariable(2, "TEST_VARIABLE_1", nil, nil)
	require.Error(t, err)
	require.Nil(t, pv)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectVariablesService_CreateVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"key": "NEW_VARIABLE",
				"value": "new value",
				"protected": false,
				"variable_type": "env_var",
				"masked": false,
				"environment_scope": "*"
			}
		`)
	})

	want := &ProjectVariable{
		Key:              "NEW_VARIABLE",
		Value:            "new value",
		VariableType:     "env_var",
		Protected:        false,
		Masked:           false,
		EnvironmentScope: "*",
	}

	pv, resp, err := client.ProjectVariables.CreateVariable(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pv)

	pv, resp, err = client.ProjectVariables.CreateVariable(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.CreateVariable(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.CreateVariable(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pv)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectVariablesService_UpdateVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/variables/NEW_VARIABLE", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBody(t, r, `{"filter":{"environment_scope":"prod"}}`)
		fmt.Fprintf(w, `
			{
				"key": "NEW_VARIABLE",
				"value": "updated value",
				"protected": false,
				"variable_type": "env_var",
				"masked": false,
				"environment_scope": "*"
			}
		`)
	})

	want := &ProjectVariable{
		Key:              "NEW_VARIABLE",
		Value:            "updated value",
		VariableType:     "env_var",
		Protected:        false,
		Masked:           false,
		EnvironmentScope: "*",
	}

	pv, resp, err := client.ProjectVariables.UpdateVariable(1, "NEW_VARIABLE", &UpdateProjectVariableOptions{Filter: &VariableFilter{EnvironmentScope: "prod"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pv)

	pv, resp, err = client.ProjectVariables.UpdateVariable(1.01, "NEW_VARIABLE", nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.UpdateVariable(1, "NEW_VARIABLE", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pv)

	pv, resp, err = client.ProjectVariables.UpdateVariable(2, "NEW_VARIABLE", nil, nil)
	require.Error(t, err)
	require.Nil(t, pv)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectVariablesService_RemoveVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/variables/VARIABLE_1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testParams(t, r, "filter%5Benvironment_scope%5D=prod")
	})

	resp, err := client.ProjectVariables.RemoveVariable(1, "VARIABLE_1", &RemoveProjectVariableOptions{Filter: &VariableFilter{EnvironmentScope: "prod"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.ProjectVariables.RemoveVariable(1.01, "VARIABLE_1", nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.ProjectVariables.RemoveVariable(1, "VARIABLE_1", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.ProjectVariables.RemoveVariable(2, "VARIABLE_1", nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
