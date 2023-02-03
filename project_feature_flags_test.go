package gitlab

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListProjectFeatureFlags(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/333/feature_flags",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			mustWriteHTTPResponse(t, w, "testdata/list_project_feature_flags.json")
		})

	actual, _, err := client.ProjectFeatureFlags.ListProjectFeatureFlags(333, &ListProjectFeatureFlagOptions{})
	if err != nil {
		t.Errorf("ProjectFeatureFlags.ListProjectFeatureFlags returned error: %v", err)
		return
	}

	createdAt1 := time.Date(2019, 11, 4, 8, 13, 51, 0, time.UTC)
	updatedAt1 := time.Date(2019, 11, 4, 8, 13, 11, 0, time.UTC)

	createdAt2 := time.Date(2019, 11, 4, 8, 13, 10, 0, time.UTC)
	updatedAt2 := time.Date(2019, 11, 4, 8, 13, 10, 0, time.UTC)

	expected := []*ProjectFeatureFlag{
		{
			Name:        "merge_train",
			Description: "This feature is about merge train",
			Active:      true,
			Version:     "new_version_flag",
			CreatedAt:   &createdAt1,
			UpdatedAt:   &updatedAt1,
			Scopes:      []*ProjectFeatureFlagScope{},
			Strategies: []*ProjectFeatureFlagStrategy{
				{
					ID:   1,
					Name: "userWithId",
					Parameters: &ProjectFeatureFlagStrategyParameter{
						UserIDs: "user1",
					},
					Scopes: []*ProjectFeatureFlagScope{
						{
							ID:               1,
							EnvironmentScope: "production",
						},
					},
				},
			},
		},
		{
			Name:        "new_live_trace",
			Description: "This is a new live trace feature",
			Active:      true,
			Version:     "new_version_flag",
			CreatedAt:   &createdAt2,
			UpdatedAt:   &updatedAt2,
			Scopes:      []*ProjectFeatureFlagScope{},
			Strategies: []*ProjectFeatureFlagStrategy{
				{
					ID:         2,
					Name:       "default",
					Parameters: &ProjectFeatureFlagStrategyParameter{},
					Scopes: []*ProjectFeatureFlagScope{
						{
							ID:               2,
							EnvironmentScope: "staging",
						},
					},
				},
			},
		},
	}

	assert.Equal(t, len(expected), len(actual))
	for i := range expected {
		assert.Equal(t, expected[i], actual[i])
	}
}

func TestGetProjectFeatureFlag(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/feature_flags/testing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_project_feature_flag.json")
	})

	actual, resp, err := client.ProjectFeatureFlags.GetProjectFeatureFlag(1, "testing")
	if err != nil {
		t.Fatalf("ProjectFeatureFlags.GetProjectFeatureFlag returned error: %v, response %v", err, resp)
	}

	date := time.Date(2020, 0o5, 13, 19, 56, 33, 0, time.UTC)
	expected := &ProjectFeatureFlag{
		Name:      "awesome_feature",
		Active:    true,
		Version:   "new_version_flag",
		CreatedAt: &date,
		UpdatedAt: &date,
		Scopes:    []*ProjectFeatureFlagScope{},
		Strategies: []*ProjectFeatureFlagStrategy{
			{
				ID:         36,
				Name:       "default",
				Parameters: &ProjectFeatureFlagStrategyParameter{},
				Scopes: []*ProjectFeatureFlagScope{
					{
						ID:               37,
						EnvironmentScope: "production",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestCreateProjectFeatureFlag(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/feature_flags/testing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		mustWriteHTTPResponse(t, w, "testdata/create_project_feature_flag.json")
	})

	actual, _, err := client.ProjectFeatureFlags.UpdateProjectFeatureFlag(1, "testing", &UpdateProjectFeatureFlagOptions{})
	if err != nil {
		t.Errorf("ProjectFeatureFlags.UpdateProjectFeatureFlag returned error: %v", err)
		return
	}

	createdAt := time.Date(2020, 5, 13, 19, 56, 33, 0, time.UTC)
	updatedAt := time.Date(2020, 5, 13, 19, 56, 33, 0, time.UTC)

	expected := &ProjectFeatureFlag{
		Name:      "awesome_feature",
		Active:    true,
		Version:   "new_version_flag",
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Scopes:    []*ProjectFeatureFlagScope{},
		Strategies: []*ProjectFeatureFlagStrategy{
			{
				ID:         36,
				Name:       "default",
				Parameters: &ProjectFeatureFlagStrategyParameter{},
				Scopes: []*ProjectFeatureFlagScope{
					{
						ID:               37,
						EnvironmentScope: "production",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestUpdateProjectFeatureFlag(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/feature_flags/testing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		mustWriteHTTPResponse(t, w, "testdata/update_project_feature_flag.json")
	})

	actual, _, err := client.ProjectFeatureFlags.UpdateProjectFeatureFlag(1, "testing", &UpdateProjectFeatureFlagOptions{})
	if err != nil {
		t.Errorf("ProjectFeatureFlags.UpdateProjectFeatureFlag returned error: %v", err)
		return
	}

	createdAt := time.Date(2020, 5, 13, 20, 10, 32, 0, time.UTC)
	updatedAt := time.Date(2020, 5, 13, 20, 10, 32, 0, time.UTC)

	expected := &ProjectFeatureFlag{
		Name:      "awesome_feature",
		Active:    true,
		Version:   "new_version_flag",
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Scopes:    []*ProjectFeatureFlagScope{},
		Strategies: []*ProjectFeatureFlagStrategy{
			{
				ID:   38,
				Name: "gradualRolloutUserId",
				Parameters: &ProjectFeatureFlagStrategyParameter{
					GroupID:    "default",
					Percentage: "25",
				},
				Scopes: []*ProjectFeatureFlagScope{
					{
						ID:               40,
						EnvironmentScope: "staging",
					},
				},
			},
			{
				ID:         37,
				Name:       "default",
				Parameters: &ProjectFeatureFlagStrategyParameter{},
				Scopes: []*ProjectFeatureFlagScope{
					{
						ID:               39,
						EnvironmentScope: "production",
					},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestDeleteProjectFeatureFlag(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/feature_flags/testing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ProjectFeatureFlags.DeleteProjectFeatureFlag(1, "testing")
	if err != nil {
		t.Errorf("ProjectFeatureFlags.DeleteProjectFeatureFlag returned error: %v", err)
		return
	}
}
