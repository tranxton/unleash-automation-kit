package main

type config struct {
	TaskNameTemplate        string `validate:"required"`
	TaskDescriptionTemplate string `validate:"required"`

	UnleashBaseURL  string `validate:"required,url"`
	UnleashApiToken string `validate:"required"`

	JiraBaseURL      string `validate:"required,url"`
	JiraProjectKey   string `validate:"required"`
	JiraIssueTypeID  string `validate:"required"`
	JiraUserEmail    string `validate:"required,email"`
	JiraUserApiToken string `validate:"required"`
}
