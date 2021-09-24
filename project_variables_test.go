package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectVariablesService_ListVariables(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/variables/TEST_VARIABLE_1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
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

	pv, resp, err := client.ProjectVariables.GetVariable(1, "TEST_VARIABLE_1", nil, nil)
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
