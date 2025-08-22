package fluxx

import (
	"net/url"
)

type Config struct {
	Instance     string
	ClientID     string
	ClientSecret string
}

func (c *Config) GetIngestrURI() string {
	// fluxx://<instance>?client_id=<client_id>&client_secret=<client_secret>
	baseURL := "fluxx://" + c.Instance
	params := url.Values{}
	params.Add("client_id", c.ClientID)
	params.Add("client_secret", c.ClientSecret)
	return baseURL + "?" + params.Encode()
}