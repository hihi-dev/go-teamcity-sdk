package teamcity_sdk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SearchBuilds_Count100(t *testing.T) {
	srv := createTestHttpServer(buildResponse)
	c := CreateGuestAuth(srv.URL)
	builds, err := c.SearchBuilds(map[string]string{
		"buildType": "myType",
		"arg2": "arg2Value",
	}, 100)
	assert.Equal(t, 2, builds.Count)
	assert.NoError(t, err)
}

func createTestHttpServer(response string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}))
	defer ts.Close()
	return ts
}

var buildResponse = `{
	"count": 2,
	"href": "/guestAuth/app/rest/builds?locator=buildType%3APhoneApp_ReleasesHotfixes,status%3ASUCCESS,defaultFilter%3Afalse,&count=100",
	"build": [{
		"id": 73950,
		"buildTypeId": "PhoneApp_ReleasesHotfixes",
		"number": "115_hotfix/HS-934--two_hover_menus_showing",
		"status": "SUCCESS",
		"state": "finished",
		"branchName": "hotfix/HS-934--two_hover_menus_showing",
		"href": "/guestAuth/app/rest/builds/id:73950",
		"webUrl": "http://teamcity.co.uk/viewLog.html?buildId=73950&buildTypeId=PhoneApp_ReleasesHotfixes"
	}, {
		"id": 62488,
		"buildTypeId": "PhoneApp_ReleasesHotfixes",
		"number": "1",
		"status": "SUCCESS",
		"state": "finished",
		"branchName": "development",
		"defaultBranch": true,
		"href": "/guestAuth/app/rest/builds/id:62488",
		"webUrl": "http://teamcity.co.uk/viewLog.html?buildId=62488&buildTypeId=PhoneApp_ReleasesHotfixes"
	}]
}`
