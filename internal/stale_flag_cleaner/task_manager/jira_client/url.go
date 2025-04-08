package jira_client

import (
	"net/url"
	"path"
)

func createIssueURL(baseURL string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "rest/api/3/issue")

	return u, nil
}
