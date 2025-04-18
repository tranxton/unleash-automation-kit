package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"unleash-automation-kit/internal/stale_flag_cleaner"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager/jira"
	jiraRepository "unleash-automation-kit/internal/stale_flag_cleaner/task_manager/jira/repository"
	"unleash-automation-kit/internal/stale_flag_cleaner/unleash"
	unleashRepository "unleash-automation-kit/internal/stale_flag_cleaner/unleash/repository"
)

func main() {
	loadEnv()
	config := loadConfigFromEnv()

	if err := validateConfig(config); err != nil {
		log.Fatalf("Invalid .env: %v", err)
	}

	cleaner := initCleaner(config)
	cleaner.CleanUpStaleFlags()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func loadConfigFromEnv() *config {
	return &config{
		TaskNameTemplate:        os.Getenv("TASK_NAME_TEMPLATE"),
		TaskDescriptionTemplate: os.Getenv("TASK_DESCRIPTION_TEMPLATE"),

		UnleashBaseURL:     os.Getenv("UNLEASH_BASE_URL"),
		UnleashProjectName: os.Getenv("UNLEASH_PROJECT_NAME"),
		UnleashApiToken:    os.Getenv("UNLEASH_PERSONAL_API_TOKEN"),

		JiraBaseURL:      os.Getenv("JIRA_BASE_URL"),
		JiraProjectKey:   os.Getenv("JIRA_PROJECT_KEY"),
		JiraIssueTypeID:  os.Getenv("JIRA_ISSUE_TYPE_ID"),
		JiraUserEmail:    os.Getenv("JIRA_USER_EMAIL"),
		JiraUserApiToken: os.Getenv("JIRA_USER_API_TOKEN"),
	}
}

func validateConfig(config *config) error {
	validate := validator.New()

	return validate.Struct(config)
}

func initCleaner(config *config) *stale_flag_cleaner.Cleaner {
	return stale_flag_cleaner.NewCleaner(
		unleash.NewUnleash(
			unleashRepository.NewRepository(
				unleashRepository.NewConfig(
					config.UnleashBaseURL,
					config.UnleashProjectName,
					config.UnleashApiToken,
				),
			),
		),
		jira.NewJira(
			jiraRepository.NewConfig(
				config.JiraBaseURL,
				config.JiraProjectKey,
				config.JiraIssueTypeID,
				config.JiraUserEmail,
				config.JiraUserApiToken,
			),
		),
		stale_flag_cleaner.NewTemplate(
			config.TaskNameTemplate,
			config.TaskDescriptionTemplate,
		),
	)
}
