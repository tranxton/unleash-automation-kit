package jira

import (
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager/jira/repository"
)

type Jira struct {
	repository *repository.Repository
}

func NewJira(repositoryConfig *repository.Config) *Jira {
	return &Jira{
		repository: repository.NewRepository(repositoryConfig),
	}
}

func (jira *Jira) FindTask(name string) (task_manager.Task, error) {
	return jira.repository.SearchIssueByName(name)
}

func (jira *Jira) CreateTask(name, description string) (task_manager.Task, error) {
	issuePayload := jira.repository.NewCreateIssueRequest(name, description)

	return jira.repository.CreateIssue(issuePayload)
}
