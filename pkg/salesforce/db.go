package salesforce

type Client struct {
	config SalesforceConfig
}

type SalesforceConfig interface {
	GetIngestrURI() string
}

func NewClient(c SalesforceConfig) (*Client, error) {
	return &Client{
		config: c,
	}, nil
}

func (c *Client) GetIngestrURI() (string, error) {
	return c.config.GetIngestrURI(), nil
}
