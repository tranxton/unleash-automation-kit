package stale_flag_cleaner

import (
	"fmt"
	"log"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager/jira"
	jiraRepository "unleash-automation-kit/internal/stale_flag_cleaner/task_manager/jira/repository"
	"unleash-automation-kit/internal/stale_flag_cleaner/unleash"
	"unleash-automation-kit/internal/stale_flag_cleaner/unleash/repository"
	unleashRepository "unleash-automation-kit/internal/stale_flag_cleaner/unleash/repository"
)

type Cleaner struct {
	unleash     *unleash.Unleash
	taskManager task_manager.TaskManager
	template    *Template
}

func NewCleaner(config *Config) *Cleaner {
	return &Cleaner{
		unleash: unleash.NewUnleash(
			unleashRepository.NewRepository(
				unleashRepository.NewConfig(
					config.UnleashBaseURL,
					config.UnleashProjectName,
					config.UnleashApiToken,
				),
			),
		),
		taskManager: jira.NewJira(
			jiraRepository.NewConfig(
				config.JiraBaseURL,
				config.JiraProjectKey,
				config.JiraIssueTypeID,
				config.JiraUserEmail,
				config.JiraUserApiToken,
			),
		),
		template: NewTemplate(
			config.TaskNameTemplate,
			config.TaskDescriptionTemplate,
		),
	}
}

func (cleaner *Cleaner) CleanUpStaleFlags() {
	features, err := cleaner.unleash.GetStaleFeatures()

	if err != nil {
		log.Fatalf("Failed to fetch stale feature flags: %v", err)
	}

	if len(features) == 0 {
		log.Println("No stale flags found")

		return
	}

	for _, feature := range features {
		message, err := cleaner.createTaskForFeature(&feature)

		if err != nil {
			log.Printf("[Feature %s] Failed to create task for: %v", feature.Name, err)

			continue
		}

		log.Println(message)
	}
}

func (cleaner *Cleaner) createTaskForFeature(feature *repository.Feature) (string, error) {
	if feature.IsTaskCreated() {
		return fmt.Sprintf("[Feature %q] Task already exists", feature.Name), nil
	}

	name := fmt.Sprintf(cleaner.template.taskName, feature.Name)
	description := fmt.Sprintf(cleaner.template.taskDescription, feature.URL)

	task, err := cleaner.taskManager.FindTask(name)
	if err != nil {
		return "", fmt.Errorf("failed to find task: %v", err)
	}

	var message string
	if task == nil {
		task, err = cleaner.taskManager.CreateTask(name, description)
		if err != nil {
			return "", fmt.Errorf("failed to create task: %v", err)
		}
		message = fmt.Sprintf("[Feature %s] Successfully created task %s", feature.Name, task.GetKey())
	} else {
		message = fmt.Sprintf("[Feature %s] Task already exists, adding tag to feature", feature.Name)
	}

	if err := cleaner.unleash.MarkTaskCreated(feature, task.GetKey()); err != nil {
		return "", fmt.Errorf("failed to tag feature: %v", err)
	}

	return message, nil
}
