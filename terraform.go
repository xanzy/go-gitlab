package gitlab

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LockInfo stores lock metadata.
//
// Only Operation and Info are required to be set by the caller of LockState.
type LockInfo struct {
	// Unique ID for the lock. This may be overridden by the lock implementation.
	// The final value of ID will be returned by the call to LockState.
	ID string `json:"ID"`

	// Terraform operation, provided by the caller.
	Operation string `json:"Operation"`

	// Extra information to store with the lock, provided by the caller.
	Info string `json:"Info"`

	// user@hostname when available
	Who string `json:"Who"`

	// Terraform version
	Version string `json:"Version"`

	// Time that the lock was taken.
	Created time.Time `json:"Created"`

	// Path to the state file when applicable.
	Path string `json:"Path"`
}

type LockData struct {
	Project           string
	StateName         string
	RequestedLockInfo LockInfo
}

// Payload is the return value from the remote state storage.
type Payload struct {
	MD5  []byte
	Data []byte
}

// TerraformService handles communication with the Terraform backend
// related methods of the GitLab API.
//
// GitLab API docs: TODO
type TerraformService struct {
	client *Client
}

func (c *TerraformService) LockState(pid interface{}, stateName string, lockInfo *LockInfo, options ...RequestOptionFunc) (*LockData, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("%s/lock", terraformStatePath(project, stateName))
	opts := []RequestOptionFunc{WithJSONBody(lockInfo), WithMD5}
	req, err := c.client.NewRequestRaw(http.MethodPost, u, append(opts, options...)...)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, nil)
	if err != nil {
		switch typedErr := err.(type) {
		case *ErrorResponse:
			switch typedErr.Response.StatusCode {
			case http.StatusForbidden:
				return nil, fmt.Errorf("remote state endpoint invalid auth")
			case http.StatusConflict, http.StatusLocked:
				existing := LockInfo{}
				err = json.Unmarshal(typedErr.Body, &existing)
				if err != nil {
					return nil, fmt.Errorf("remote state already locked, failed to unmarshal body: %v", err)
				}
				return nil, fmt.Errorf("remote state already locked: ID=%s", existing.ID)
			default:
				return nil, fmt.Errorf("unexpected HTTP response code %d", typedErr.Response.StatusCode)
			}
		default:
			return nil, err
		}
	}
	return &LockData{
		Project:           project,
		StateName:         stateName,
		RequestedLockInfo: *lockInfo,
	}, nil
}

func (c *TerraformService) UnlockState(lockData *LockData, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("%s/lock", terraformStatePath(lockData.Project, lockData.StateName))

	// NB: lock id is NOT used. Just like in the Terraform HTTP backend
	opts := []RequestOptionFunc{WithJSONBody(&lockData.RequestedLockInfo), WithMD5}
	req, err := c.client.NewRequestRaw(http.MethodPost, u, append(opts, options...)...)
	if err != nil {
		return err
	}
	_, err = c.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *TerraformService) GetState(pid interface{}, stateName string, options ...RequestOptionFunc) (*Payload, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := terraformStatePath(project, stateName)
	req, err := c.client.NewRequestRaw(http.MethodGet, u, options...)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	resp, err := c.client.Do(req, &buf)
	if err != nil {
		switch typedErr := err.(type) {
		case *ErrorResponse:
			switch typedErr.Response.StatusCode {
			case http.StatusNotFound:
				return nil, nil
			default:
				return nil, fmt.Errorf("unexpected HTTP response code: %d", typedErr.Response.StatusCode)
			}
		default:
			return nil, err
		}
	}
	// TODO it's annoying to check the response code in two places.
	// TODO Implement a DoRaw() method that does less magic and allows for more control?
	switch resp.StatusCode {
	case http.StatusOK:
		// Handled after
	case http.StatusNoContent:
		return nil, nil
	default:
		return nil, fmt.Errorf("unexpected HTTP response code: %d", resp.StatusCode)
	}

	data := buf.Bytes()

	// If there was no data, then return nil
	if len(data) == 0 {
		return nil, nil
	}

	// Check for the MD5
	var MD5 []byte
	if raw := resp.Header.Get("Content-Md5"); raw != "" {
		decodedMD5, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to decode Content-Md5 '%s': %v", raw, err)
		}
		MD5 = decodedMD5
	} else {
		// Generate the MD5
		hash := md5.Sum(data)
		MD5 = hash[:]
	}
	return &Payload{
		MD5:  MD5,
		Data: data,
	}, nil
}

func (c *TerraformService) PutState(lockData *LockData, data []byte, options ...RequestOptionFunc) error {
	u := terraformStatePath(lockData.Project, lockData.StateName)
	queryParams := struct {
		ID string `url:"ID"`
	}{
		ID: lockData.RequestedLockInfo.ID,
	}
	opts := []RequestOptionFunc{
		WithBody(data, "application/json"),
		WithMD5,
		WithQueryParameters(&queryParams),
	}
	req, err := c.client.NewRequestRaw(http.MethodPost, u, append(opts, options...)...)
	if err != nil {
		return err
	}
	_, err = c.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *TerraformService) DeleteState(pid interface{}, stateName string, options ...RequestOptionFunc) error {
	project, err := parseID(pid)
	if err != nil {
		return err
	}
	u := terraformStatePath(project, stateName)
	req, err := c.client.NewRequestRaw(http.MethodDelete, u, options...)
	if err != nil {
		return err
	}
	_, err = c.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func terraformStatePath(project, stateName string) string {
	return fmt.Sprintf("projects/%s/terraform/state/%s", pathEscape(project), pathEscape(stateName))
}
