package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInstanceVariablesService_ListVariables(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"key": "TEST_VARIABLE_1",
					"variable_type": "env_var",
					"value": "TEST_1",
					"protected": false,
					"masked": false,
					"raw": true
				}
			]
		`)
	})

	want := []*InstanceVariable{{
		Key:          "TEST_VARIABLE_1",
		Value:        "TEST_1",
		VariableType: "env_var",
		Protected:    false,
		Masked:       false,
		Raw:          true,
	}}

	ivs, resp, err := client.InstanceVariables.ListVariables(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ivs)

	ivs, resp, err = client.InstanceVariables.ListVariables(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ivs)
}

func TestInstanceVariablesService_ListVariables_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	ivs, resp, err := client.InstanceVariables.ListVariables(nil)
	require.Error(t, err)
	require.Nil(t, ivs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceVariablesService_GetVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables/TEST_VARIABLE_1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
				"key": "TEST_VARIABLE_1",
				"variable_type": "env_var",
				"value": "TEST_1",
				"protected": false,
				"masked": false,
				"raw": true
			}
		`)
	})

	want := &InstanceVariable{
		Key:          "TEST_VARIABLE_1",
		Value:        "TEST_1",
		VariableType: "env_var",
		Protected:    false,
		Masked:       false,
		Raw:          true,
	}

	iv, resp, err := client.InstanceVariables.GetVariable("TEST_VARIABLE_1", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, iv)

	iv, resp, err = client.InstanceVariables.GetVariable("TEST_VARIABLE_1", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, iv)

	iv, resp, err = client.InstanceVariables.GetVariable("TEST_VARIABLE_2", nil)
	require.Error(t, err)
	require.Nil(t, iv)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceVariablesService_CreateVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"key": "NEW_VARIABLE",
				"value": "new value",
				"variable_type": "env_var",
				"protected": false,
				"masked": false,
				"raw": true
			}
		`)
	})

	want := &InstanceVariable{
		Key:          "NEW_VARIABLE",
		Value:        "new value",
		VariableType: "env_var",
		Protected:    false,
		Masked:       false,
		Raw:          true,
	}

	iv, resp, err := client.InstanceVariables.CreateVariable(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, iv)

	iv, resp, err = client.InstanceVariables.CreateVariable(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, iv)
}

func TestInstanceVariablesService_StatusInternalServerError(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusInternalServerError)
	})

	iv, resp, err := client.InstanceVariables.CreateVariable(nil)
	require.Error(t, err)
	require.Nil(t, iv)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestInstanceVariablesService_UpdateVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables/NEW_VARIABLE", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
				"key": "NEW_VARIABLE",
				"value": "updated value",
				"variable_type": "env_var",
				"protected": false,
				"masked": false,
				"raw": true
			}
		`)
	})

	want := &InstanceVariable{
		Key:          "NEW_VARIABLE",
		Value:        "updated value",
		VariableType: "env_var",
		Protected:    false,
		Masked:       false,
		Raw:          true,
	}

	iv, resp, err := client.InstanceVariables.UpdateVariable("NEW_VARIABLE", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, iv)

	iv, resp, err = client.InstanceVariables.UpdateVariable("NEW_VARIABLE", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, iv)

	iv, resp, err = client.InstanceVariables.UpdateVariable("NEW_VARIABLE_1", nil)
	require.Error(t, err)
	require.Nil(t, iv)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceVariablesService_RemoveVariable(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/ci/variables/NEW_VARIABLE", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.InstanceVariables.RemoveVariable("NEW_VARIABLE", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.InstanceVariables.RemoveVariable("NEW_VARIABLE", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.InstanceVariables.RemoveVariable("NEW_VARIABLE_1", nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
