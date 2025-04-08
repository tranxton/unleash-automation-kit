package unleash

import (
	"bytes"
	"encoding/json"
)

func (unleash *Unleash) MarkTaskCreated(feature *Feature, taskID string) error {
	return unleash.addTag(feature, newTag("deleteTaskCreated", taskID))
}

func (unleash *Unleash) addTag(feature *Feature, tag *Tag) error {
	body, _ := json.Marshal(tag)
	URL, _ := addFeatureTagURL(unleash.Config.baseURL, feature.Name)

	_, err := unleash.doRequest("POST", URL.String(), bytes.NewReader(body))

	if err != nil {
		return err
	}

	return nil
}
