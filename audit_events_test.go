package gitlab

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTargetID_UnmarshalJSON Ensures that the json unmarshals as expected with mixed float/string like values
func TestTargetID_UnmarshalJSON(t *testing.T) {
	jsonIn := []byte(`[
		{
			"id": 1,
			"details": {
				"target_id": "project/path/here"
			}
		},
		{
			"id": 2,
			"details": {
				"target_id": 123
			}
		}]`)

	var auditEvents []*AuditEvent
	err := json.Unmarshal(jsonIn, &auditEvents)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, TargetID("project/path/here"), auditEvents[0].Details.TargetID)
	assert.Equal(t, TargetID("123"), auditEvents[1].Details.TargetID)
}
