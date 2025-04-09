package repository

type Tag struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewTag(typeName, value string) *Tag {
	return &Tag{
		Type:  typeName,
		Value: value,
	}
}
