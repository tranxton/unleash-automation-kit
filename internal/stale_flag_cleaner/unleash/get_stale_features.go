package unleash

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type featuresResponse struct {
	Features []Feature `json:"features"`
}

func (unleash *Unleash) GetStaleFeatures() ([]Feature, error) {
	url := fmt.Sprintf(searchFeatureURL, unleash.Config.baseURL, "state=IS:stale")

	responseBody, err := unleash.doRequest("GET", url, nil)
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
