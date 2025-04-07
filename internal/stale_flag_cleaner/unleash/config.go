package unleash

type Config struct {
	baseURL          string
	personalAPIToken string
}

func NewConfig(baseURL, personalAPIToken string) *Config {
	return &Config{
		baseURL:          baseURL,
		personalAPIToken: personalAPIToken,
	}
}
