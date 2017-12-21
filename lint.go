package gitlab

type LintService struct {
	client *Client
}

type LintResponse struct {
	Status string   `json:"status"`
	Errors []string `json:"errors"`
}

// Lint validates .gitlab-ci content
//
// GitLab API docs: https://docs.gitlab.com/ce/api/lint.html
func (l *LintService) Lint(content string, options ...OptionFunc) (*LintResponse, *Response, error) {
	opts := struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	req, err := l.client.NewRequest("POST", "ci/lint", opts, options)
	if err != nil {
		return nil, nil, err
	}

	var lint LintResponse
	resp, err := l.client.Do(req, &lint)
	if err != nil {
		return nil, resp, err
	}
	return &lint, resp, nil
}
