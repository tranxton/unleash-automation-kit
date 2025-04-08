package unleash

import (
	"fmt"
	"io"
	"net/http"
)

type Unleash struct {
	Config *Config
	Client *http.Client
}

func NewUnleash(config *Config) *Unleash {
	return &Unleash{
		Config: config,
		Client: http.DefaultClient,
	}
}

func (unleash *Unleash) doRequest(method, URL string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, URL, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", unleash.Config.personalAPIToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := unleash.Client.Do(req)
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
