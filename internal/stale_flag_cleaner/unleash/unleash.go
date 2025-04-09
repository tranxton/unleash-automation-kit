package unleash

import (
	"unleash-automation-kit/internal/stale_flag_cleaner/unleash/repository"
)

type Unleash struct {
	repository *repository.Repository
}

func NewUnleash(repository *repository.Repository) *Unleash {
	return &Unleash{
		repository: repository,
	}
}
func (unleash *Unleash) MarkTaskCreated(feature *repository.Feature, taskID string) error {
	return unleash.repository.AddTagToFeature(feature, repository.NewTag("deleteTaskCreated", taskID))
}

func (unleash *Unleash) GetStaleFeatures() ([]repository.Feature, error) {
	return unleash.repository.SearchStaleFeatures()
}
