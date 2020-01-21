package config

/*
A helper class to parameters from "options" configuration section.

Example:
 config := NewConfigParamsFromTuples(
    ...
    "options.param1", "ABC",
    "options.param2", 123
);

 options := OptionsResolver.resolve(config); // Result: param1=ABC;param2=123
*/

type TOptionsResolver struct{}

var OptionsResolver *TOptionsResolver = &TOptionsResolver{}

// Resolves configuration section from component configuration parameters.
// Parameters:
// 			 - config: ConfigParams
// 			  configuration parameters
// Returns *ConfigParams
// configuration parameters from "options" section
func (c *TOptionsResolver) Resolve(config *ConfigParams) *ConfigParams {
	var options = config.GetSection("options")
	return options
}

// Resolves an "options" configuration section from component configuration parameters.
// Parameters:
// 			 - config: ConfigParams
// 				configuration parameters
// 			 - configAsDefault: boolean
//  		When set true the method returns the entire parameter set when "options" section is not found. Default: false
// Returns *ConfigParams
// configuration parameters from "options" section
func (c *TOptionsResolver) ResolveWithDefault(config *ConfigParams) *ConfigParams {
	var options = c.Resolve(config)

	if options.Len() == 0 {
		options = config
	}

	return options
}
