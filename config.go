package skyscanner

// Config contains application secrets
type Config struct {
	MashapeKey *string `yaml:"mashape_key"`
	BaseURL    *string `yaml:"base_url"`
}

var apiConfig Config

// SetConfig sets global variables for use in the SDK
func SetConfig(config *Config) {
	apiConfig = *config
}
