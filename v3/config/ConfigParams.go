package config

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/data"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/reflect"
)

/*
Contains a key-value map with configuration parameters. All values stored as strings and can be serialized
as JSON or string forms. When retrieved the values can be automatically converted on read using GetAsXXX methods.
The keys are case-sensitive, so it is recommended to use consistent C-style as: "my_param"

Configuration parameters can be broken into sections and subsections using dot notation as:
"section1.subsection1.param1". Using GetSection method all parameters from specified section can be extracted from a ConfigMap.

The ConfigParams supports serialization from/to plain strings as: "key1=123;key2=ABC;key3=2016-09-16T00:00:00.00Z"

ConfigParams are used to pass configurations to IConfigurable objects. They also serve as a basis for more concrete configurations such as ConnectionParams or CredentialParams (in the Pip.Services components package).

see
IConfigurable

see
StringValueMap

Example:
 config := NewConfigParamsFromTuples(
    "section1.key1", "AAA",
    "section1.key2", 123,
    "section2.key1", true
);

config.GetAsString("section1.key1"); // Result: AAA
config.GetAsInteger("section1.key1"); // Result: 0

section1 = config.GetSection("section2");
section1.GetAsString("key1"); // Result: true
*/

type ConfigParams struct {
	data.StringValueMap
}

// Creates a new empty ConfigParams object.
// Returns *ConfigParams
// a new empty ConfigParams object.
func NewEmptyConfigParams() *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewEmptyStringValueMap(),
	}
}

// Creates a new ConfigParams from map.
// Parameters:
// 			- values ...map[string]string
// Returns *ConfigParams
// a newly created ConfigParams.
func NewConfigParams(values map[string]string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMap(values),
	}
}

// Creates a new ConfigParams object filled with key-value pairs from specified object.
// Parameters:
// 			- value interface{}
// 			an object with key-value pairs used to initialize a new ConfigParams.
// Returns *ConfigParams
// a new ConfigParams object.
func NewConfigParamsFromValue(value interface{}) *ConfigParams {
	values := reflect.RecursiveObjectReader.GetProperties(value)
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromValue(values),
	}
}

// Creates a new ConfigParams object filled with provided key-value pairs called tuples.
//Tuples parameters contain a sequence of key1, value1, key2, value2, ... pairs.
// see
// StringValueMap.fromTuplesArray
// Parameters:
// 			- tuples ...interface{}
// the tuples to fill a new ConfigParams object.

// Returns ConfigParams
// a new ConfigParams object.
func NewConfigParamsFromTuples(tuples ...interface{}) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromTuplesArray(tuples),
	}
}

// Creates a new StringValueMap from a list of key-value pairs called tuples.
// The method is similar to fromTuples but tuples are passed as array instead of parameters.
// Parameters:
// 			 - tuples []interface{}
// 			a list of values where odd elements are keys and the following even elements are values
// Returns *ConfigParams
// a newly created ConfigParams.
func NewConfigParamsFromTuplesArray(tuples []interface{}) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromTuplesArray(tuples),
	}
}

// Creates a new ConfigParams object filled with key-value pairs serialized as a string.
// see
// StringValueMap.fromString
// Parameters:
// 				- line: string
// 				a string with serialized key-value pairs as "key1=value1;key2=value2;..."
//				Example: "Key1=123;Key2=ABC;Key3=2016-09-16T00:00:00.00Z"
// Returns *ConfigParams
// a new ConfigParams object.
func NewConfigParamsFromString(line string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromString(line),
	}
}

// Creates a new ConfigParams by merging two or more maps.
// Maps defined later in the list override values from previously defined maps.
// Parameters:
// 			  - maps ...map[string]string
// an array of maps to be merged
// Returns *ConfigParams
// a newly created ConfigParams.
func NewConfigParamsFromMaps(maps ...map[string]string) *ConfigParams {
	return &ConfigParams{
		StringValueMap: *data.NewStringValueMapFromMaps(maps...),
	}
}

// Gets a list with all 1st level section names.
// Returns []string
// a list of section names stored in this ConfigMap.

func (c *ConfigParams) GetSectionNames() []string {
	sections := []string{}

	for key := range c.Value() {
		pos := strings.Index(key, ".")
		section := key
		if pos > 0 {
			section = key[0:pos]
		}

		// Perform case sensitive search
		found := false
		for index := 0; index < len(sections); index++ {
			if section == sections[index] {
				found = true
				break
			}
		}

		if !found {
			sections = append(sections, section)
		}
	}

	return sections
}

// Gets parameters from specific section stored in this ConfigMap. The section name is removed from parameter keys.
// Parameters:
// 			- section: string
// 			name of the section to retrieve configuration parameters from.
// Returns *ConfigParams
// all configuration parameters that belong to the section named 'section'.
func (c *ConfigParams) GetSection(section string) *ConfigParams {
	result := NewEmptyConfigParams()
	prefix := section + "."

	for key := range c.Value() {
		// Prevents exception on the next line
		if len(key) < len(prefix) {
			continue
		}

		// Perform case sensitive match
		keyPrefix := key[0:len(prefix)]
		if keyPrefix == prefix {
			sectionKey := key[len(prefix):]
			result.Put(sectionKey, c.Get(key))
		}
	}

	return result
}

// Adds parameters into this ConfigParams under specified section.
// Keys for the new parameters are appended with section dot prefix.
// Parameters:
// 			- section: string
// 			name of the section where add new parameters
// 			- sectionParams: *ConfigParams
// 			new parameters to be added.
func (c *ConfigParams) AddSection(section string, sectionParams *ConfigParams) {
	if section == "" {
		panic("Section name cannot be empty")
	}

	if sectionParams != nil {
		for key := range sectionParams.Value() {
			sectionKey := key

			if len(sectionKey) > 0 {
				sectionKey = section + "." + sectionKey
			} else {
				sectionKey = section
			}

			value := (*sectionParams).Get(key)

			c.Put(sectionKey, value)
		}
	}
}

// Overrides parameters with new values from specified ConfigParams and returns a new ConfigParams object.
// see
// setDefaults
// Parameters:
//  		- configParams: *ConfigParams
// 			ConfigMap with parameters to override the current values.
// Returns *ConfigParams
// a new ConfigParams object.
func (c *ConfigParams) Override(configParams *ConfigParams) *ConfigParams {
	return NewConfigParamsFromMaps(c.Value(), configParams.Value())
}

// Set default values from specified ConfigParams and returns a new ConfigParams object.
// see
// override
// Parameters:
// 		- defaultConfigParams: *ConfigParams
// 			ConfigMap with default parameter values.
// Returns *ConfigParams
// a new ConfigParams object.
func (c *ConfigParams) SetDefaults(defaults *ConfigParams) *ConfigParams {
	return NewConfigParamsFromMaps(defaults.Value(), c.Value())
}
