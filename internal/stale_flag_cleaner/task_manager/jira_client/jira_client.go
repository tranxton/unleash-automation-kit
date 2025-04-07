package jira_client

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Jira struct {
	config *Config
	Client *http.Client
}

func NewJira(config *Config) *Jira {
	return &Jira{
		config: config,
		Client: http.DefaultClient,
	}
}

func (jira *Jira) doRequest(method, url string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", jira.config.userEmail, jira.config.userApiToken)))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := jira.Client.Do(req)
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
