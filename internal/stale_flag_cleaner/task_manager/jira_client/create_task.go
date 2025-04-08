package jira_client

import (
	"bytes"
	"encoding/json"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager"
)

type createIssuePayload struct {
	Fields issueFields `json:"fields"`
}

type issueFields struct {
	Project     project     `json:"project"`
	Summary     string      `json:"summary"`
	Description description `json:"description"`
	IssueType   issueType   `json:"issuetype"`
}

type project struct {
	Key string `json:"key"`
}

type issueType struct {
	ID string `json:"id"`
}

type description struct {
	Type    string    `json:"type"`
	Version int       `json:"version"`
	Content []content `json:"content"`
}

type content struct {
	Type    string         `json:"type"`
	Content []textFragment `json:"content"`
}

type textFragment struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (jira *Jira) CreateTask(name, description string) (task_manager.Task, error) {
	task := jira.newCreateIssuePayload(name, description)
	body, _ := json.Marshal(task)
	URL, _ := createIssueURL(jira.config.baseURL)

	responseBody, err := jira.doRequest("POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var issue Issue
	if err := json.Unmarshal(responseBody, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (jira *Jira) newCreateIssuePayload(summary, descriptionContent string) *createIssuePayload {
	return &createIssuePayload{
		Fields: issueFields{
			Project: project{
				Key: jira.config.projectKey,
			},
			Summary: summary,
			Description: description{
				Type:    "doc",
				Version: 1,
				Content: []content{
					{
						Type: "paragraph",
						Content: []textFragment{
							{
								Type: "text",
								Text: descriptionContent,
							},
						},
					},
				},
			},
			IssueType: issueType{
				ID: jira.config.issueTypeID,
			},
		},
	}
}
