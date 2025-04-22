package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"unleash-automation-kit/internal/stale_flag_cleaner"
)

func main() {
	loadEnv()
	config := loadConfigFromEnv()

	if err := validateConfig(config); err != nil {
		log.Fatalf("Invalid .env: %v", err)
	}

	cleaner := stale_flag_cleaner.NewCleaner(config)
	cleaner.CleanUpStaleFlags()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func loadConfigFromEnv() *stale_flag_cleaner.Config {
	return &stale_flag_cleaner.Config{
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

func validateConfig(config *stale_flag_cleaner.Config) error {
	validate := validator.New()

	return validate.Struct(config)
}
