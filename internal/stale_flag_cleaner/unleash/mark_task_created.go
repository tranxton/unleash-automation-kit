package unleash

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (unleash *Unleash) MarkTaskCreated(feature *Feature, taskID string) error {
	return unleash.addTag(feature, newTag("deleteTaskCreated", taskID))
}

func (unleash *Unleash) addTag(feature *Feature, tag *Tag) error {
	body, _ := json.Marshal(tag)

	_, err := unleash.doRequest("POST", fmt.Sprintf(addFeatureTagURL, unleash.Config.baseURL, feature.Name), bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}
