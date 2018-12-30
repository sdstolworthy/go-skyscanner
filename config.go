package skyscanner

// Config contains application secrets
type Config struct {
	MashapeKey *string `yaml:"mashape_key"`
	BaseURL    *string `yaml:"base_url"`
}

var apiConfig Config

func setConfig(config *Config) {
	apiConfig = *config
}
