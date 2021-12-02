package gitlab

import "net/http"

// MarkdownService handles communication with the markdown endpoint of the gitlab API
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/markdown.html
type MarkdownService struct {
	client *Client
}

// Markdown represents gitlab markdown
type Markdown struct {
	HTML string `json:"html"`
}

// MarkdownOptions represents the available Markdown() options.
type MarkdownOptions struct {
	Text                    string `json:"text,omitempty"`
	GitlabFlavouredMarkdown bool   `json:"gfm,omitempty"`
	Project                 string `json:"project,omitempty"`
}

// Markdown creates html from the given markdown
//
// Gitlab API docs:
func (s *MarkdownService) Markdown(opt *MarkdownOptions, options ...RequestOptionFunc) (*Markdown, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "markdown", opt, options)
	if err != nil {
		return nil, nil, err
	}

	markdown := new(Markdown)
	response, err := s.client.Do(req, markdown)
	if err != nil {
		return nil, response, err
	}
	return markdown, response, nil
}
