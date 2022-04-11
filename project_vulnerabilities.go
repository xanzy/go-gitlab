package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

type ProjectVulnerabilitiesService struct {
	client *Client
}

type ProjectVulnerability struct {
	AuthorID                int       `json:"author_id"`
	Confidence              string    `json:"confidence"`
	CreatedAt               time.Time `json:"created_at"`
	Description             string    `json:"description"`
	DismissedAt             time.Time `json:"dismissed_at"`
	DismissedByID           int       `json:"dismissed_by_id"`
	DueDate                 time.Time `json:"due_date"`
	Finding                 Finding   `json:"finding"`
	ID                      int       `json:"id"`
	LastEditedAt            time.Time `json:"last_edited_at"`
	LastEditedByID          int       `json:"last_edited_by_id"`
	Project                 Project   `json:"project"`
	ProjectDefaultBranch    string    `json:"project_default_branch"`
	ReportType              string    `json:"report_type"`
	ResolvedAt              time.Time `json:"resolved_at"`
	ResolvedByID            int       `json:"resolved_by_id"`
	ResolvedOnDefaultBranch bool      `json:"resolved_on_default_branch"`
	Severity                string    `json:"severity"`
	StartDate               time.Time `json:"start_date"`
	State                   string    `json:"state"`
	Title                   string    `json:"title"`
	UpdatedAt               time.Time `json:"updated_at"`
	UpdatedByID             int       `json:"updated_by_id"`
}

type Finding struct {
	Confidence          string    `json:"confidence"`
	CreatedAt           time.Time `json:"created_at"`
	ID                  int       `json:"id"`
	LocationFingerprint string    `json:"location_fingerprint"`
	MetadataVersion     string    `json:"metadata_version"`
	Name                string    `json:"name"`
	PrimaryIdentifierID int       `json:"primary_identifier_id"`
	ProjectFingerprint  string    `json:"project_fingerprint"`
	ProjectID           int       `json:"project_id"`
	RawMetadata         string    `json:"raw_metadata"`
	ReportType          string    `json:"report_type"`
	ScannerID           int       `json:"scanner_id"`
	Severity            string    `json:"severity"`
	UpdatedAt           time.Time `json:"updated_at"`
	UUID                string    `json:"uuid"`
	VulnerabilityID     int       `json:"vulnerability_id"`
}

type ListProjectVulnerabilitiesOptions struct {
	ListOptions
}

func (s *ProjectVulnerabilitiesService) ListProjectVulnerabilities(pid interface{}, opt *ListProjectVulnerabilitiesOptions, options ...RequestOptionFunc) ([]*ProjectVulnerability, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/vulnerabilities", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectVulnerability
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
