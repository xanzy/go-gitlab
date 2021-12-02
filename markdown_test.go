package gitlab

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestMarkdown(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	const htmlResponse = "<h1>Testing</h1>"

	mux.HandleFunc("/api/v4/markdown", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodPost)
		writer.WriteHeader(http.StatusOK)
		markdown := Markdown{HTML: htmlResponse}
		resp, _ := json.Marshal(markdown)
		_, _ = writer.Write(resp)
	})

	opt := &MarkdownOptions{
		Text:                    "# Testing",
		GitlabFlavouredMarkdown: true,
		Project:                 "some/sub/group/project",
	}
	markdown, resp, err := client.Markdown.Markdown(opt)
	if err != nil {
		t.Fatalf("GetMarkdown returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GetMarkdown retruned status expected %q but was %q", http.StatusOK, resp.Status)
	}

	if markdown.HTML != htmlResponse {
		t.Fatalf("GetMarkdown returned wrong response, expected HTML to be %q but was %q", htmlResponse, markdown.HTML)
	}
}
