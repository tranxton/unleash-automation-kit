package repository

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

func searchIssueURL(baseURL, JQLQuery string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "rest/api/3/search/jql")

	q := u.Query()
	q.Set("jql", JQLQuery)
	q.Set("fields", "id,key")
	u.RawQuery = q.Encode()

	return u, nil
}
