package teamcity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) PostCommentOnBuild(buildId int64, comment string) error {
	url := fmt.Sprintf("/builds/id:%d/comment", buildId)
	c.headers["Content-Type"] = "text/plain"
	c.headers["Origin"] = c.baseUrl
	r, err := c.doRequestWithPrefix("PUT", url, strings.NewReader(comment))
	if err != nil {
		return err
	}
	if 204 != r.StatusCode {
		return http.ErrBodyNotAllowed
	}
	return nil
}

type UserComment struct {
	Timestamp string `json:"timestamp"`
	Text      string `json:"text"`
}

type UserCommentResponse struct {
	Comment UserComment `json:"comment"`
}

func (c *Client) GetCommentOnBuild(buildId int64) (*UserCommentResponse, error) {
	url := fmt.Sprintf("/builds/id:%d?fields=comment", buildId)
	r, err := c.doRequestWithPrefix("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp := &UserCommentResponse{}
	b, be := ioutil.ReadAll(r.Body)
	if be != nil {
		return nil, be
	}
	jsErr :=json.Unmarshal(b, resp)
	defer r.Body.Close()
	if jsErr != nil {
		return nil, jsErr
	}
	return resp, nil
}

func (c *Client) DeleteCommentOnBuild(buildId int64) error {
	url := fmt.Sprintf("/builds/id:%d/comment", buildId)
	r, err := c.doRequestWithPrefix("DELETE", url, nil)
	if err != nil {
		return err
	}
	if 204 != r.StatusCode {
		return http.ErrBodyNotAllowed
	}
	return nil
}
