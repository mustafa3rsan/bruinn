package salesforce

import "net/url"

type Config struct {
	Instance string
	ClientId string
	ClientSecret string
}

func (c *Config) GetIngestrURI() string {
	v := url.Values{}
	v.Set("instance", c.Instance)
	v.Set("client_id", c.ClientId)
	v.Set("client_secret", c.ClientSecret)
	return "salesforce://" + c.Instance + "?" + v.Encode()
}
