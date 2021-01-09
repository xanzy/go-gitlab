package gitlab

// AuthMethodFunc is used to configure the auth method for a GitLab API client.
type AuthMethodFunc func(*Client) error

// BasicAuth configures the Gitlab client to use a username and password.
func BasicAuth(username, password string) AuthMethodFunc {
	return func(c *Client) error {
		c.authType = basicAuth
		c.username = username
		c.password = password
		return nil
	}
}

// DeployTokenAuth configures the Gitlab client to use a deploy token.
func DeployTokenAuth(token string) AuthMethodFunc {
	return func(c *Client) error {
		c.authType = deployToken
		c.token = token
		return nil
	}
}

// JobTokenAuth configures the Gitlab client to use a job token.
func JobTokenAuth(token string) AuthMethodFunc {
	return func(c *Client) error {
		c.authType = jobToken
		c.token = token
		return nil
	}
}

// OAuthTokenAuth configures the Gitlab client to use an oauth token.
func OAuthTokenAuth(token string) AuthMethodFunc {
	return func(c *Client) error {
		c.authType = oAuthToken
		c.token = token
		return nil
	}
}

// PrivateTokenAuth configures the Gitlab client to use a private token.
func PrivateTokenAuth(token string) AuthMethodFunc {
	return func(c *Client) error {
		c.authType = privateToken
		c.token = token
		return nil
	}
}
