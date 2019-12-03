package teamcity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *Client) GetArtifactsForBuildId(buildId int64) (*BuildArtifactsList, error) {
	url := fmt.Sprintf("/builds/id:%d/artifacts", buildId)
	resp, err := c.doRequest("GET", url)
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
