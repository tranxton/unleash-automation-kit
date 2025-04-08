package unleash

import (
	"bytes"
	"encoding/json"
)

type featuresResponse struct {
	Features []Feature `json:"features"`
}

func (unleash *Unleash) GetStaleFeatures() ([]Feature, error) {
	URL, _ := getStaleFeaturesURL(unleash.Config.baseURL, unleash.Config.projectName)

	responseBody, err := unleash.doRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	var decodeResponse featuresResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decodeResponse); err != nil {
		return nil, err
	}

	for i := range decodeResponse.Features {
		decodeResponse.Features[i].setUrl(unleash.Config.baseURL)
	}

	return decodeResponse.Features, nil
}
