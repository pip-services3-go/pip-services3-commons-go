package config

type TOptionsResolver struct{}

var OptionsResolver *TOptionsResolver = &TOptionsResolver{}

func (c *TOptionsResolver) Resolve(config *ConfigParams) *ConfigParams {
	var options = config.GetSection("options")
	return options
}

func (c *TOptionsResolver) ResolveWithDefault(config *ConfigParams) *ConfigParams {
	var options = c.Resolve(config)

	if options.Len() == 0 {
		options = config
	}

	return options
}
