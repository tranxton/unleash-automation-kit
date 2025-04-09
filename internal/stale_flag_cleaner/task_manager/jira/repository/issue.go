package repository

type Issue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

func (issue *Issue) GetKey() string {
	return issue.Key
}
