package teamcity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *Client) GetArtifactsForBuildId(buildId int64) (*BuildArtifactsList, error) {
	url := fmt.Sprintf("/builds/id:%d/artifacts", buildId)
	resp, err := c.doRequestWithPrefix("GET", url, nil)
	if err != nil {
		return nil, err
	}
	art := &BuildArtifactsList{}
	js, be := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if be != nil {
		return nil, be
	}
	je := json.Unmarshal(js, art)
	return art, je
}

// TeamCity requires turning the received 'href' of an artifact into a usable download path
func (c *Client) ResolveArtifactUrl(href string) (string, error) {
	jsn, err := c.doRequest("GET", c.baseUrl + href, nil)
	if err != nil {
		return "", err
	}
	resp := &buildArtifactHrefResponse{}
	bd, be := ioutil.ReadAll(jsn.Body)
	defer jsn.Body.Close()
	if be != nil {
		return "", be
	}
	je := json.Unmarshal(bd, resp)
	return c.baseUrl + resp.Content.Href, je
}

type buildArtifactHrefResponse struct {
	Name             string      `json:"name"`
	Size             int64       `json:"size"`
	ModificationTime string      `json:"modificationTime"`
	Href             string      `json:"href"`
	Content          hrefContent `json:"content"`
	Children         hrefContent `json:"children"`
}

type hrefContent struct {
	Href string `json:"href"`
}

type BuildArtifactsList struct {
	Artifacts []BuildArtifact `json:"file"`
	Count     int             `json:"count"`
}

type BuildArtifact struct {
	Name             string `json:"name"`
	Size             int64  `json:"size"`
	ModificationTime string `json:"modificationTime"`
	HREF             string `json:"href"`
}
