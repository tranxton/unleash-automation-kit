package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repository struct {
	config *Config
	client *http.Client
}

func NewRepository(config *Config) *Repository {
	return &Repository{
		config: config,
		client: http.DefaultClient,
	}
}

func (repository *Repository) AddTagToFeature(feature *Feature, tag *Tag) error {
	body, _ := json.Marshal(tag)
	URL, _ := addTagToFeatureURL(repository.config.baseURL, feature.Name)

	_, err := repository.doRequest("POST", URL.String(), bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) SearchStaleFeatures() ([]Feature, error) {
	URL, _ := searchStaleFeaturesURL(repository.config.baseURL, repository.config.projectName)

	responseBody, err := repository.doRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	var decodeResponse FeaturesResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&decodeResponse); err != nil {
		return nil, err
	}

	for i := range decodeResponse.Features {
		decodeResponse.Features[i].setUrl(repository.config.baseURL)
	}

	return decodeResponse.Features, nil
}

func (repository *Repository) doRequest(method, URL string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, URL, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", repository.config.personalAPIToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := repository.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("request execution error: %s, %s", resp.Status, responseBody)
	}

	return responseBody, nil
}
