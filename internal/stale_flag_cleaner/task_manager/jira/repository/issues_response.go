package repository

type IssuesResponse struct {
	Issues []*Issue `json:"issues"`
}
