package config

import "strings"

type TNameResolver struct{}

var NameResolver *TNameResolver = &TNameResolver{}

func (c *TNameResolver) Resolve(config *ConfigParams) string {
	return c.ResolveWithDefault(config, "")
}

func (c *TNameResolver) ResolveWithDefault(config *ConfigParams, defaultName string) string {
	var name = config.GetAsString("name")

	if name == "" {
		name = config.GetAsString("id")
	}

	if name == "" {
		var descriptorStr = config.GetAsString("descriptor")
		if descriptorStr != "" {
			tokens := strings.Split(descriptorStr, ":")
			if len(tokens) == 5 {
				name = tokens[3]
			}
		}
	}

	if name == "" {
		name = defaultName
	}

	return name
}
