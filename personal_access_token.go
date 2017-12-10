package gitlab

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	gitlabAPITimeout        = 10 * time.Second
	pathRoot                = "/"
	pathSignIn              = "/users/sign_in"
	pathPersonalAccessToken = "/profile/personal_access_tokens"
)

type tokenParams struct {
	endpointURL *url.URL  // Gitlab endpoint
	login       string    // Gitlab user login used for HTTP basic auth
	password    string    // Gitlab user password used for HTTP basic auth
	name        string    // token name
	expiresAt   time.Time // token expiration date
}

type csrfParams struct {
	param string
	value string
}

// NewPersonalAccessTokenFromBaseAuth uses HTTP basic auth in order to get
// new personal access token for Gitlab API
func NewPersonalAccessTokenFromBaseAuth(
	endpoint, login, password string,
	tokenName string, tokenExpiresAt time.Time,
) (string, error) {

	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	p := &tokenParams{
		endpointURL: endpointURL,
		login:       login,
		password:    password,
		name:        tokenName,
		expiresAt:   tokenExpiresAt,
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return "", nil
	}

	client := &http.Client{
		Timeout: gitlabAPITimeout,
		Jar:     cookieJar,
	}

	csrf1, err := obtainRootCSRFToken(client, p)
	if err != nil {
		return "", err
	}
	csrf2, err := obtainSignInCSRFToken(client, p, csrf1)
	if err != nil {
		return "", err
	}

	return obtainPersonalAccessToken(client, p, csrf2)
}

// obtainRootCSRFToken requests main page of Gitlab in order to obtain
// CSRF token for a further use
func obtainRootCSRFToken(client *http.Client, p *tokenParams) (*csrfParams, error) {
	targetURL := *p.endpointURL
	targetURL.Path = path.Join(targetURL.Path, pathRoot)

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code when obtaining root CSRF token: %v", resp.StatusCode)
	}

	return scrapeCSRFToken(resp.Body)
}

// obtainSignInCSRFToken signs into Gitlab with provided credentials in order
// to obtain CSRF token for a further use
func obtainSignInCSRFToken(client *http.Client, p *tokenParams, csrf *csrfParams) (*csrfParams, error) {
	targetURL := *p.endpointURL
	targetURL.Path = path.Join(targetURL.Path, pathSignIn)

	v := url.Values{}
	v.Set("user[login]", p.login)
	v.Set("user[password]", p.password)
	v.Set("user[remember_me]", "0")
	v.Set("utf8", "✓")
	v.Set(csrf.param, csrf.value)

	form := strings.NewReader(v.Encode())

	req, err := http.NewRequest("POST", targetURL.String(), form)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code when obtaining sign in page CSRF token: %v", resp.StatusCode)
	}

	return scrapeCSRFToken(resp.Body)
}

func obtainPersonalAccessToken(client *http.Client, p *tokenParams, csrf *csrfParams) (string, error) {
	targetURL := *p.endpointURL
	targetURL.Path = path.Join(targetURL.Path, pathPersonalAccessToken)

	v := url.Values{}
	v.Set("personal_access_token[expires_at]", p.expiresAt.Format("2006-01-02"))
	v.Set("personal_access_token[name]", p.name)
	v.Set("personal_access_token[scopes][]", "api")
	v.Set("utf8", "✓")
	v.Set(csrf.param, csrf.value)

	form := strings.NewReader(v.Encode())

	req, err := http.NewRequest("POST", targetURL.String(), form)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code when obtaining personal access token: %v", resp.StatusCode)
	}

	return scrapePersonalAccessToken(resp.Body)
}

// scrapeCSRFToken parses web-page in search of CSRF tokens
func scrapeCSRFToken(body io.Reader) (*csrfParams, error) {
	root, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	csrfParamMatcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Meta {
			return scrape.Attr(n, "name") == "csrf-param"
		}
		return false
	}
	csrfTokenMatcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Meta {
			return scrape.Attr(n, "name") == "csrf-token"
		}
		return false
	}

	csrfParamNode, ok := scrape.Find(root, csrfParamMatcher)
	if !ok {
		return nil, fmt.Errorf("Can't find csrf-param attribute")
	}
	csrfTokenNode, ok := scrape.Find(root, csrfTokenMatcher)
	if !ok {
		return nil, fmt.Errorf("Can't find csrf-token attribute")
	}

	result := &csrfParams{
		param: scrape.Attr(csrfParamNode, "content"),
		value: scrape.Attr(csrfTokenNode, "content"),
	}
	return result, nil
}

func scrapePersonalAccessToken(body io.Reader) (string, error) {
	root, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Input {
			return scrape.Attr(n, "name") == "created-personal-access-token"
		}
		return false
	}

	data, ok := scrape.Find(root, matcher)
	if !ok {
		return "", fmt.Errorf("Can't find created-personal-access-token attribute")
	}
	return scrape.Attr(data, "value"), nil
}
