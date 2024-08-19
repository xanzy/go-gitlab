package gitlab

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dispatcherContextKey struct{}

func newDispatcherContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, dispatcherContextKey{}, "DispatcherContext")
}

func testDispatcherContext(ctx context.Context, t *testing.T) {
	v, ok := ctx.Value(dispatcherContextKey{}).(string)
	assert.True(t, ok)
	assert.Equal(t, "DispatcherContext", v)
}

func TestDispatcher_Dispatch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dispatcher := NewDispatcher()
		dispatcher.RegisterListeners(&testListener{t: t})

		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/webhook", r.URL.Path)
		assert.NoError(t,
			dispatcher.DispatchRequest(r,
				DispatchRequestWithContext(newDispatcherContext(r.Context())),
			),
		)

		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer srv.Close()

	tests := []struct {
		name      string
		eventType EventType
		body      []byte
	}{
		{"build", EventTypeBuild, loadFixture("testdata/webhooks/build.json")},
		{"commit comment", EventTypeNote, loadFixture("testdata/webhooks/note_commit.json")},
		{"deployment", EventTypeDeployment, loadFixture("testdata/webhooks/deployment.json")},
		{"feature flag", EventTypeFeatureFlag, loadFixture("testdata/webhooks/feature_flag.json")},
		{"group resource access token", EventTypeResourceAccessToken, loadFixture("testdata/webhooks/resource_access_token_group.json")},
		{"issue comment", EventTypeNote, loadFixture("testdata/webhooks/note_issue.json")},
		{"issue", EventTypeIssue, loadFixture("testdata/webhooks/issue.json")},
		{"job", EventTypeJob, loadFixture("testdata/webhooks/job.json")},
		{"member", EventTypeMember, loadFixture("testdata/webhooks/member.json")},
		{"merge comment", EventTypeNote, loadFixture("testdata/webhooks/note_merge_request.json")},
		{"merge", EventTypeMergeRequest, loadFixture("testdata/webhooks/merge_request.json")},
		{"pipeline", EventTypePipeline, loadFixture("testdata/webhooks/pipeline.json")},
		{"push", EventTypePush, loadFixture("testdata/webhooks/push.json")},
		{"release", EventTypeRelease, loadFixture("testdata/webhooks/release.json")},
		{"snippet comment", EventTypeNote, loadFixture("testdata/webhooks/note_snippet.json")},
		{"subgroup", EventTypeSubGroup, loadFixture("testdata/webhooks/subgroup.json")},
		{"tag", EventTypeTagPush, loadFixture("testdata/webhooks/tag_push.json")},
		{"wiki page", EventTypeWikiPage, loadFixture("testdata/webhooks/wiki_page.json")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, srv.URL+"/webhook", bytes.NewReader(tt.body))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Gitlab-Event", string(tt.eventType))

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var buf bytes.Buffer
			_, _ = buf.ReadFrom(resp.Body)
			assert.Equal(t, `{"status":"ok"}`, buf.String())
		})
	}
}

type testListener struct {
	t *testing.T
}

var (
	_ BuildListener                      = (*testListener)(nil)
	_ CommitCommentListener              = (*testListener)(nil)
	_ DeploymentListener                 = (*testListener)(nil)
	_ FeatureFlagListener                = (*testListener)(nil)
	_ GroupResourceAccessTokenListener   = (*testListener)(nil)
	_ IssueCommentListener               = (*testListener)(nil)
	_ IssueListener                      = (*testListener)(nil)
	_ JobListener                        = (*testListener)(nil)
	_ MemberListener                     = (*testListener)(nil)
	_ MergeCommentListener               = (*testListener)(nil)
	_ MergeListener                      = (*testListener)(nil)
	_ PipelineListener                   = (*testListener)(nil)
	_ ProjectResourceAccessTokenListener = (*testListener)(nil)
	_ PushListener                       = (*testListener)(nil)
	_ ReleaseListener                    = (*testListener)(nil)
	_ SnippetCommentListener             = (*testListener)(nil)
	_ SubGroupListener                   = (*testListener)(nil)
	_ TagListener                        = (*testListener)(nil)
	_ WikiPageListener                   = (*testListener)(nil)
)

func (t *testListener) OnBuild(ctx context.Context, event *BuildEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "gitlab-org/gitlab-test", event.ProjectName)
	return nil
}

func (t *testListener) OnCommitComment(ctx context.Context, event *CommitCommentEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnDeployment(ctx context.Context, event *DeploymentEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "test-deployment-webhooks", event.Project.Name)
	return nil
}

func (t *testListener) OnFeatureFlag(ctx context.Context, event *FeatureFlagEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "gitlabhq/gitlab-test", event.Project.PathWithNamespace)
	return nil
}

func (t *testListener) OnGroupResourceAccessToken(ctx context.Context, event *GroupResourceAccessTokenEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "expiring_access_token", event.EventName)
	return nil
}

func (t *testListener) OnIssueComment(ctx context.Context, event *IssueCommentEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnIssue(ctx context.Context, event *IssueEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnJob(ctx context.Context, event *JobEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "auto_deploy:start", event.BuildName)
	return nil
}

func (t *testListener) OnMember(ctx context.Context, event *MemberEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "User1", event.UserName)
	return nil
}

func (t *testListener) OnMergeComment(ctx context.Context, event *MergeCommentEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil

}

func (t *testListener) OnMerge(ctx context.Context, event *MergeEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnPipeline(ctx context.Context, event *PipelineEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnProjectResourceAccessToken(ctx context.Context, event *ProjectResourceAccessTokenEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "expiring_access_token", event.EventName)
	return nil
}

func (t *testListener) OnPush(ctx context.Context, event *PushEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "mike/diaspora", event.Project.PathWithNamespace)
	return nil
}

func (t *testListener) OnRelease(ctx context.Context, event *ReleaseEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Project Name", event.Project.Name)
	return nil
}

func (t *testListener) OnSnippetComment(ctx context.Context, event *SnippetCommentEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Gitlab Test", event.Project.Name)
	return nil
}

func (t *testListener) OnSubGroup(ctx context.Context, event *SubGroupEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "SubGroup 1", event.Name)
	return nil
}

func (t *testListener) OnTag(ctx context.Context, event *TagEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "Example", event.Project.Name)
	return nil
}

func (t *testListener) OnWikiPage(ctx context.Context, event *WikiPageEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, "awesome-project", event.Project.Name)
	return nil
}
