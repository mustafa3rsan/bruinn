
type FluxxConnection struct {
	Name string `yaml:"name,omitempty" json:"name" mapstructure:"name"`
	Instance string `yaml:"instance,omitempty" json:"instance" mapstructure:"instance"`
	ClientId string `yaml:"client_id,omitempty" json:"client_id" mapstructure:"client_id"`
	ClientSecret string `yaml:"client_secret,omitempty" json:"client_secret" mapstructure:"client_secret"`
}

func (c FluxxConnection) GetName() string {
	return c.Name
}
