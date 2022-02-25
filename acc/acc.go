package acc

import (
	"os"
	"testing"
	"math/rand"
	"time"

	"github.com/xanzy/go-gitlab"
)

type Config struct {
	BaseURL string
	Token string
}

func setup(t *testing.T) (*gitlab.Client, Config) {
	config := newConfig()
	client := newTestClient(t, config)
	return client, config
}

func newConfig() Config {
	baseURL := os.Getenv("GITLAB_BASE_URL")
	token := os.Getenv("GITLAB_TOKEN")

	config := Config{BaseURL: baseURL, Token: token}
	return config
}

func newTestClient(t *testing.T, config Config) *gitlab.Client {
	client, err := gitlab.NewOAuthClient(config.Token, gitlab.WithBaseURL(config.BaseURL))
	if err != nil {
		t.Errorf("failed to instantiate new GitLab OAuth client: %v", err)
	}
	return client
}

func init() {
	rand.Seed(time.Now().UnixNano())
}