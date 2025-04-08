package unleash

type Feature struct {
	Name    string `json:"name"`
	Project string `json:"project"`
	Tags    []Tag  `json:"tags"`

	URL string `json:"-"`
}

func (feature *Feature) setUrl(baseURL string) {
	URL, _ := getFeatureURL(baseURL, feature.Project, feature.Name)

	feature.URL = URL.String()
}

func (feature *Feature) IsTaskCreated() bool {
	return feature.hasTag("deleteTaskCreated")
}

func (feature *Feature) hasTag(name string) bool {
	for _, tag := range feature.Tags {
		if tag.Type == name {
			return true
		}
	}

	return false
}
