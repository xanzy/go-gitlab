package gitlab

import (
	"encoding/json"
	"net/http"
	"testing"
)

const markdownHTMLResponse = "<h1>Testing</h1>"

func TestRender(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/markdown", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodPost)
		writer.WriteHeader(http.StatusOK)
		markdown := Markdown{HTML: markdownHTMLResponse}
		resp, _ := json.Marshal(markdown)
		_, _ = writer.Write(resp)
	})

	opt := &RenderOptions{
		Text:                    String("# Testing"),
		GitlabFlavouredMarkdown: Bool(true),
		Project:                 String("some/sub/group/project"),
	}
	markdown, resp, err := client.Markdown.Render(opt)
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Render returned status, expected %q but got %q", http.StatusOK, resp.Status)
	}

	if markdown.HTML != markdownHTMLResponse {
		t.Fatalf("Render returned wrong response, expected %q but got %q",
			markdownHTMLResponse, markdown.HTML)
	}
}
