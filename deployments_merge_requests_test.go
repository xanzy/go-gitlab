package gitlab

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeploymentMergeRequestsService_ListDeploymentMergeRequests(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/278964/deployments/2/merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "assignee_id=Any&with_labels_details=true&with_merge_status_recheck=true")
		mustWriteHTTPResponse(t, w, "testdata/get_merge_requests.json")
	})

	opts := ListMergeRequestsOptions{
		AssigneeID:             AssigneeID(UserIDAny),
		WithLabelsDetails:      Bool(true),
		WithMergeStatusRecheck: Bool(true),
	}

	mergeRequests, _, err := client.DeploymentMergeRequests.ListDeploymentMergeRequests(278964, 2, &opts)
	require.NoError(t, err)
	require.Equal(t, 20, len(mergeRequests))

	validStates := []string{"opened", "closed", "locked", "merged"}
	detailedMergeStatuses := []string{
		"blocked_status",
		"broken_status",
		"checking",
		"ci_must_pass",
		"ci_still_running",
		"discussions_not_resolved",
		"draft_status",
		"external_status_checks",
		"mergeable",
		"not_approved",
		"not_open",
		"policies_denied",
		"unchecked",
	}
	allCreatedBefore := time.Date(2019, 8, 21, 0, 0, 0, 0, time.UTC)
	allCreatedAfter := time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC)

	for _, mr := range mergeRequests {
		require.Equal(t, 278964, mr.ProjectID)
		require.Contains(t, validStates, mr.State)
		assert.Less(t, mr.CreatedAt.Unix(), allCreatedBefore.Unix())
		assert.Greater(t, mr.CreatedAt.Unix(), allCreatedAfter.Unix())
		assert.LessOrEqual(t, mr.CreatedAt.Unix(), mr.UpdatedAt.Unix())
		assert.LessOrEqual(t, mr.TaskCompletionStatus.CompletedCount, mr.TaskCompletionStatus.Count)
		require.Contains(t, detailedMergeStatuses, mr.DetailedMergeStatus)

		// list requests do not provide these fields:
		assert.Nil(t, mr.Pipeline)
		assert.Nil(t, mr.HeadPipeline)
		assert.Equal(t, "", mr.DiffRefs.HeadSha)
	}
}
