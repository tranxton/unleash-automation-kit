package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager"
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

func (repository *Repository) CreateIssue(createIssueRequest *CreateIssueRequest) (task_manager.Task, error) {
	body, _ := json.Marshal(createIssueRequest)
	URL, _ := createIssueURL(repository.config.baseURL)

	responseBody, err := repository.doRequest("POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var issue Issue
	if err := json.Unmarshal(responseBody, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (repository *Repository) SearchIssueByName(name string) (task_manager.Task, error) {
	JQLQuery := fmt.Sprintf("summary~\"%s\" AND project = \"%s\"", name, repository.config.projectKey)
	URL, _ := searchIssueURL(repository.config.baseURL, JQLQuery)

	responseBody, err := repository.doRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	var issueResponse IssuesResponse
	if err = json.Unmarshal(responseBody, &issueResponse); err != nil {
		return nil, err
	}

	if len(issueResponse.Issues) == 0 {
		return nil, nil
	}

	return issueResponse.Issues[0], nil
}

func (repository *Repository) doRequest(method, URL string, requestBody io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, URL, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", repository.makeAuthToken()))
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

func (repository *Repository) makeAuthToken() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", repository.config.userEmail, repository.config.userApiToken)))
}
