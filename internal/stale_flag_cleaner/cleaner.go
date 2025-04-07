package stale_flag_cleaner

import (
	"fmt"
	"log"
	"unleash-automation-kit/internal/stale_flag_cleaner/task_manager"
	"unleash-automation-kit/internal/stale_flag_cleaner/unleash"
)

type Cleaner struct {
	unleash     *unleash.Unleash
	taskManager task_manager.TaskManager
	template    *Template
}

func NewCleaner(unleash *unleash.Unleash, task task_manager.TaskManager, template *Template) *Cleaner {
	return &Cleaner{
		unleash:     unleash,
		taskManager: task,
		template:    template,
	}
}

func (cleaner *Cleaner) CleanUpStaleFlags() {
	features, err := cleaner.unleash.GetStaleFeatures()

	if err != nil {
		log.Fatalf("Failed to fetch stale feature flags: %v", err)
	}

	for _, feature := range features {
		task, err := cleaner.createTaskForFeature(&feature)

		if err != nil {
			log.Printf("Failed to create task for feature %q: %v", feature.Name, err)

			continue
		}

		log.Printf("Successfully created task %s for feature %q", feature.Name, task.GetKey())
	}
}

func (cleaner *Cleaner) createTaskForFeature(feature *unleash.Feature) (task_manager.Task, error) {
	if feature.IsTaskCreated() {
		return nil, fmt.Errorf("task already exist")
	}

	name := fmt.Sprintf(cleaner.template.taskName, feature.Name)
	description := fmt.Sprintf(cleaner.template.taskDescription, feature.URL)
	task, err := cleaner.taskManager.CreateTask(name, description)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %v", err)
	}

	if err := cleaner.unleash.MarkTaskCreated(feature, task.GetKey()); err != nil {
		return nil, fmt.Errorf("failed to tag feature: %v", err)
	}

	return task, nil
}
