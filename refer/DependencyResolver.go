package refer

import (
	conf "github.com/pip-services3-go/pip-services3-commons-go/config"
	convert "github.com/pip-services3-go/pip-services3-commons-go/convert"
)

/*
Helper class for resolving component dependencies.

The resolver is configured to resolve named dependencies by specific locator. During deployment the dependency locator can be changed.

This mechanism can be used to clarify specific dependency among several alternatives. Typically components are configured to retrieve the first dependency that matches logical group, type and version. But if container contains more than one instance and resolution has to be specific about those instances, they can be given a unique name and dependency resolvers can be reconfigured to retrieve dependencies by their name.

Configuration parameters
dependencies:

[dependency name 1]: Dependency 1 locator (descriptor)
...
[dependency name N]: Dependency N locator (descriptor)
References
References must match configured dependencies.

Example:

*/

type DependencyResolver struct {
	dependencies map[string]interface{}
	references   IReferences
}

// Creates a new instance of the dependency resolver.
// Returns *DependencyResolver
func NewDependencyResolver() *DependencyResolver {
	return &DependencyResolver{
		dependencies: map[string]interface{}{},
		references:   nil,
	}
}

// Creates a new instance of the dependency resolver.
// see
// ConfigParams
// see
// configure
// see
// IReferences
// see
// setReferences
// Parameters:
// 			- config ConfigParams
// 			 default configuration where key is dependency name and value is locator (descriptor)
// 			- references IReferences
// 			 default component references
// Returns *DependencyResolver
func NewDependencyResolverWithParams(config *conf.ConfigParams, references IReferences) *DependencyResolver {
	c := NewDependencyResolver()

	if config != nil {
		c.Configure(config)
	}

	if references != nil {
		c.SetReferences(references)
	}

	return c
}

//Configures the component with specified parameters.
// see
// ConfigParams
// Parameters:
// 			- config *conf.ConfigParams
// 			configuration parameters to set.
func (c *DependencyResolver) Configure(config *conf.ConfigParams) {
	dependencies := config.GetSection("dependencies")
	names := dependencies.Keys()
	for _, name := range names {
		locator := dependencies.Get(name)
		if locator == "" {
			continue
		}

		descriptor, err := ParseDescriptorFromString(locator)
		if err == nil {
			c.dependencies[name] = descriptor
		} else {
			c.dependencies[name] = locator
		}
	}
}

// Sets the component references. References must match configured dependencies.
// Parameters:
// 			- references IReferences
// 			references to set.
func (c *DependencyResolver) SetReferences(references IReferences) {
	c.references = references
}

// Adds a new dependency into this resolver.
// Parameters:
// 			- name string
// 			the dependency's name.
// 			locator interface{}
// 			the locator to find the dependency by.

func (c *DependencyResolver) Put(name string, locator interface{}) {
	c.dependencies[name] = locator
}

// Locate dependency by name
// Parameters:
//			- name string
// 			dependency name
// Returns interface{}
// located dependency
func (c *DependencyResolver) Locate(name string) interface{} {
	if name == "" {
		panic("Dependency name cannot be empty")
	}

	if c.references == nil {
		panic("References shall be set")
	}

	return c.dependencies[name]
}

// Gets all optional dependencies by their name.
// Parameters:
// 			- name string
// the dependency name to locate.
// Returns []interface{}
// a list with found dependencies or empty list of no dependencies was found.
func (c *DependencyResolver) GetOptional(name string) []interface{} {
	locator := c.Locate(name)
	if locator == nil {
		return []interface{}{}
	}
	return c.references.GetOptional(locator)
}

// Gets all required dependencies by their name. At least one dependency must be present.
// If no dependencies was found it throws a ReferenceError
// throws
// a ReferenceError if no dependencies were found.
// Parameters:
// 			- name string
// 			the dependency name to locate.
// Returns []interface{}
// a list with found dependencies.
func (c *DependencyResolver) GetRequired(name string) ([]interface{}, error) {
	locator := c.Locate(name)
	if locator == nil {
		err := NewReferenceError("", name)
		return []interface{}{}, err
	}

	return c.references.GetRequired(locator)
}

// Gets one optional dependency by its name.
// Parameters:
// 			- name string
// 			the dependency name to locate.
// Returns interface{}
// a dependency reference or nil of the dependency was not found
func (c *DependencyResolver) GetOneOptional(name string) interface{} {
	locator := c.Locate(name)
	if locator == nil {
		return nil
	}
	return c.references.GetOneOptional(locator)
}

// Gets one required dependency by its name. At least one dependency must present.
// If the dependency was found it throws a ReferenceError
// throws
// a ReferenceError if dependency was not found.
// Parameters:
// 			- name string
// 			the dependency name to locate.
// Returns interface {}, error
// a dependency reference and error

func (c *DependencyResolver) GetOneRequired(name string) (interface{}, error) {
	locator := c.Locate(name)
	if locator == nil {
		err := NewReferenceError("", name)
		return nil, err
	}
	return c.references.GetOneRequired(locator)
}

// Finds all matching dependencies by their name.
// throws
// a ReferenceError of required is true and no dependencies found.
// Parameters:
// 			- name string
// 			the dependency name to locate.
// 			- required bool
// 			true to raise an exception when no dependencies are found.
// Returns []interface{}, error
// a list of found dependencies and error
func (c *DependencyResolver) Find(name string, required bool) ([]interface{}, error) {
	if name == "" {
		panic("Name cannot be empty")
	}

	locator := c.Locate(name)
	if locator == nil {
		if required {
			err := NewReferenceError("", name)
			return []interface{}{}, err
		}
		return []interface{}{}, nil
	}

	return c.references.Find(locator, required)
}

// Creates a new DependencyResolver from a list of key-value pairs
// called tuples where key is dependency name and value the depedency locator (descriptor).
// see
// NewDependencyResolverFromTuplesArray
// Parameters:
// 			- tuples ...interface{}
// 			a list of values where odd elements are dependency name and
//			the following even elements are dependency locator (descriptor)
// Returns *DependencyResolver
// a newly created DependencyResolver.
func NewDependencyResolverFromTuples(tuples ...interface{}) *DependencyResolver {
	result := NewDependencyResolver()
	if len(tuples) == 0 {
		return result
	}

	for index := 0; index < len(tuples); index += 2 {
		if index+1 >= len(tuples) {
			break
		}

		name := convert.StringConverter.ToString(tuples[index])
		locator := tuples[index+1]

		result.Put(name, locator)
	}

	return result
}
