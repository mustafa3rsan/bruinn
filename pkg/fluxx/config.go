package fluxx

import "fmt"

type Config struct {
	Instance     string
	ClientID     string
	ClientSecret string
}

func (c *Config) GetIngestrURI() string {
	return fmt.Sprintf("fluxx://%s?client_id=%s&client_secret=%s", c.Instance, c.ClientID, c.ClientSecret)
}