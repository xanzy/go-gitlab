package gitlab

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"

	"github.com/google/go-querystring/query"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// RequestOptionFunc can be passed to all API requests to customize the API request.
type RequestOptionFunc func(*retryablehttp.Request) error

// WithSudo takes either a username or user ID and sets the SUDO request header
func WithSudo(uid interface{}) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		user, err := parseID(uid)
		if err != nil {
			return err
		}
		req.Header.Set("SUDO", user)
		return nil
	}
}

// WithContext runs the request with the provided context
func WithContext(ctx context.Context) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		*req = *req.WithContext(ctx)
		return nil
	}
}

// WithBody sets the request body. Also sets the Content-Type header (if non-empty).
func WithBody(body interface{}, contentType string) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		// Cannot mutate just the body of the existing req retryablehttp.Request because it's in a private field.
		// Have to resort to tricks.
		newReq, err := retryablehttp.NewRequest(req.Method, req.URL.String(), body)
		if err != nil {
			return err
		}
		newContentLength := newReq.ContentLength        // Save the new content length
		*newReq.Request = *req.Request                  // Overwrite new http.Request with the old one to preserve any possible earlier mutations
		newReq.Request.ContentLength = newContentLength // Restore new content length
		*req = *newReq                                  // Set private field(s) too
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}

		return nil
	}
}

// WithJSONBody marshals bodyObj into JSON representation and sets the request body.
func WithJSONBody(bodyObj interface{}) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		jsonBody, err := json.Marshal(bodyObj)
		if err != nil {
			return err
		}
		return WithBody(jsonBody, "application/json")(req)
	}
}

// WithMD5 sets the Content-Md5 header by calculating MD5 hash of the request body.
func WithMD5(req *retryablehttp.Request) error {
	// TODO this can be optimized once https://github.com/hashicorp/go-retryablehttp/pull/88 is merged
	// req.WriteTo(md5 hasher)
	bodyBytes, err := req.BodyBytes()
	if err != nil {
		return err
	}
	md5Sum := md5.Sum(bodyBytes)
	req.Header.Set("Content-Md5", base64.StdEncoding.EncodeToString(md5Sum[:]))
	return nil
}

// WithQueryParameters sets the query parameters.
func WithQueryParameters(queryParams interface{}) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		q, err := query.Values(queryParams)
		if err != nil {
			return err
		}
		req.URL.RawQuery = q.Encode()
		return nil
	}
}
