package config

/*
An interface to set configuration parameters to an object.

It can be added to any existing class by implementing a single configure() method.

If you need to emphasis the fact that configure() method can be called multiple
times to change object configuration in runtime, use IReconfigurable interface instead.
*/
type IConfigurable interface {
	// 	Configures object by passing configuration parameters.

	// Parameters:
	//  - config: ConfigParams
	//  configuration parameters to be set.
	Configure(config *ConfigParams)
}
