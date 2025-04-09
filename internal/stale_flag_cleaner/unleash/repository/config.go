package repository

type Config struct {
	baseURL          string
	projectName      string
	personalAPIToken string
}

func NewConfig(baseURL, projectName, personalAPIToken string) *Config {
	return &Config{
		baseURL:          baseURL,
		projectName:      projectName,
		personalAPIToken: personalAPIToken,
	}
}
