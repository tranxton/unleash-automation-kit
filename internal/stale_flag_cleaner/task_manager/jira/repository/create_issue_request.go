package repository

type CreateIssueRequest struct {
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

func (repository *Repository) NewCreateIssueRequest(summary, descriptionContent string) *CreateIssueRequest {
	return &CreateIssueRequest{
		Fields: issueFields{
			Project: project{
				Key: repository.config.projectKey,
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
				ID: repository.config.issueTypeID,
			},
		},
	}
}
