package acc

import (
	"testing"
)


func TestGitlabAuth_withAccessToken(t *testing.T) {
	client, _ := setup(t)

	_, _, err := client.Users.CurrentUser()
	if err != nil {
		t.Errorf("failed to query current user, something went wrong with the client instantition: %v", err)
	}
}