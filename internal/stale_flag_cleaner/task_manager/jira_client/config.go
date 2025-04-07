package jira_client

type Config struct {
	baseURL      string
	projectKey   string
	issueTypeID  string
	userEmail    string
	userApiToken string
}

func NewConfig(baseURL, projectKey, issueTypeID, userEmail, userApiToken string) *Config {
	return &Config{
		baseURL:      baseURL,
		projectKey:   projectKey,
		issueTypeID:  issueTypeID,
		userEmail:    userEmail,
		userApiToken: userApiToken,
	}
}
