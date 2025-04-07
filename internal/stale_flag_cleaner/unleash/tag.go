package unleash

type Tag struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func newTag(typeName, value string) *Tag {
	return &Tag{
		Type:  typeName,
		Value: value,
	}
}
