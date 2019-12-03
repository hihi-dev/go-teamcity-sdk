package teamcity_sdk

import (
	"fmt"
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
func (c *Client) doRequest(path string) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	client := &http.Client{}
	url := c.createBasePath() + path
	log.Println("Requesting: ", url)
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}

// The path to the rest API
func (c *Client) createBasePath() string {
	return fmt.Sprintf("%s/%s/app/rest", c.baseUrl, c.authType)
}
