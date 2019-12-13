package teamcity

import (
	"fmt"
	"io"
	"log"
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
func (c *Client) doRequest(method string, fullUrl string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	client := &http.Client{}
	log.Println("teamcity-sdk Request:", method, fullUrl)
	req, _ := http.NewRequest(method, fullUrl, body)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}

func (c *Client) doRequestWithPrefix(method string, path string, body io.Reader) (*http.Response, error) {
	return c.doRequest(method, c.createBasePath() + path, body)
}

// The path to the rest API
func (c *Client) createBasePath() string {
	return fmt.Sprintf("%s/%s/app/rest", c.baseUrl, c.authType)
}
