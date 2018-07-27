package config

type IConfigurable interface {
	Configure(config *ConfigParams)
}
