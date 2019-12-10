package teamcity

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	authType string
	baseUrl  string
	headers  map[string]string
}

func CreateGuestAuth(url string) *Client {
	return &Client{
		authType: "guestAuth",
		baseUrl:  url,
		headers:  map[string]string{},
	}
}

// Perform an action on the API against this path
func (c *Client) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	client := &http.Client{}
	req, _ := http.NewRequest(method, c.createBasePath() + path, nil)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}

// The path to the rest API
func (c *Client) createBasePath() string {
	return fmt.Sprintf("%s/%s/app/rest", c.baseUrl, c.authType)
}
