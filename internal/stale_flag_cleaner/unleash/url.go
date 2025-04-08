package unleash

import (
	"fmt"
	"net/url"
	"path"
)

func getFeatureURL(baseURL, project, feature string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "projects", project, "features", feature)

	return u, nil
}

func getStaleFeaturesURL(baseURL, project string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "api/admin/search/features")

	q := u.Query()
	q.Set("project", fmt.Sprintf("IS:%s", project))
	q.Set("state", "IS:stale")
	u.RawQuery = q.Encode()

	return u, nil
}

func addFeatureTagURL(baseURL, feature string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "api/admin/features", feature, "tags")

	return u, nil
}
