package gitlab

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

// TerraformService handles communication with the Terraform backend
// related methods of the GitLab API.
//
// GitLab API docs: TODO add a link that describes the API
type TerraformService struct {
	client *Client
}

// TerraformState is the return value from the remote state storage and
// contains both the state as well as an MD5 hash of the state.
//
// GitLab API docs: TODO add a link that describes the fields
type TerraformState struct {
	MD5  string          `json:"md5"`
	Data json.RawMessage `json:"data"`
}

// LockInfo stores lock metadata. Only operation and info are required
// to be set by the caller of LockState.
//
// GitLab API docs: TODO add a link that describes the fields
type LockInfo struct {
	ID        string    `url:"id" json:"id"`
	Operation string    `url:"-" json:"operation"`
	Info      string    `url:"-" json:"info"`
	Who       string    `url:"-" json:"who"`
	Version   string    `url:"-" json:"version"`
	Created   time.Time `url:"-" json:"created"`
	Path      string    `url:"-" json:"path"`
}

func (c *TerraformService) LockState(pid interface{}, name string, lockInfo *LockInfo, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/terraform/state/%s/lock", pathEscape(project), pathEscape(name))

	req, err := c.client.NewRequest("POST", u, lockInfo, options)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}

func (c *TerraformService) UnlockState(pid interface{}, name string, lockInfo *LockInfo, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/terraform/state/%s/unlock", pathEscape(project), pathEscape(name))

	req, err := c.client.NewRequest("POST", u, lockInfo, options)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}

func (c *TerraformService) GetState(pid interface{}, name string, options ...RequestOptionFunc) (*TerraformState, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/terraform/state/%s", pathEscape(project), pathEscape(name))

	req, err := c.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ts := new(TerraformState)
	resp, err := c.client.Do(req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, err
}

type PutStateOptions struct {
	MD5  *string         `url:"md5,omitempty" json:"md5,omitempty"`
	Data json.RawMessage `url:"data,omitempty" json:"data,omitempty"`
}

func (c *TerraformService) PutState(pid interface{}, name string, lockInfo *LockInfo, opt *PutStateOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/terraform/state/%s", pathEscape(project), pathEscape(name))

	req, err := c.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, err
	}

	// Add the lock ID as query param.
	q, err := query.Values(lockInfo)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = q.Encode()

	// If still required the MD5 header could still be added.
	req.Header.Set("Content-Md5", *opt.MD5)

	// Or if the body remains to be the bare Terraform state,
	// the custom MD5 header could also be added like this.
	body, err := req.BodyBytes()
	if err != nil {
		return nil, err
	}
	md5Sum := md5.Sum(body)
	req.Header.Set("Content-Md5", base64.StdEncoding.EncodeToString(md5Sum[:]))

	return c.client.Do(req, nil)
}

func (c *TerraformService) DeleteState(pid interface{}, name string, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/terraform/state/%s", pathEscape(project), pathEscape(name))

	req, err := c.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req, nil)
}
