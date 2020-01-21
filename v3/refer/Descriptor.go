package refer

import (
	"strings"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/errors"
)

/*
Locator type that most often used in PipServices toolkit. It locates components using several fields:

Group: a package or just named group of components like "pip-services"
Type: logical component type that defines it's contract like "persistence"
Kind: physical implementation type like "mongodb"
Name: unique component name like "default"
Version: version of the component contract like "1.0"
The locator matching can be done by all or only few selected fields. The fields that shall be excluded from the matching must be set to "*" or null. That approach allows to implement many interesting scenarios. For instance:

Locate all loggers (match by type and version)
Locate persistence components for a microservice (match by group and type)
Locate specific component by its name (match by name)
Example
locator1 := NewDescriptor("mygroup", "connector", "aws", "default", "1.0");
locator2 := NewDescriptorFromString("mygroup:connector:*:*:1.0");

locator1.Match(locator2);        // Result: true
locator1.Equal(locator2);        // Result: true
locator1.ExactMatch(locator2);    // Result: false
*/

type Descriptor struct {
	group   string
	typ     string
	kind    string
	name    string
	version string
}

// Creates a new instance of the descriptor.
// Parameters:
// 			- group string
// 			a logical component group
// 			- type string
// 			a logical component type or contract
// 			- kind string
// 			a component implementation type
// 			- name string
// 			a unique component name
// 			- version string
// 			a component implementation version
// Returns *Descriptor

func NewDescriptor(group string, typ string, kind string, name string, version string) *Descriptor {
	if "*" == group {
		group = ""
	}
	if "*" == typ {
		typ = ""
	}
	if "*" == kind {
		kind = ""
	}
	if "*" == name {
		name = ""
	}
	if "*" == version {
		version = ""
	}

	return &Descriptor{group: group, typ: typ, kind: kind, name: name, version: version}
}

// Gets the component's logical group.
// Returns string
// the component's logical group
func (c *Descriptor) Group() string {
	return c.group
}

// Gets the component's logical type.
// Returns string
// the component's logical type.
func (c *Descriptor) Type() string {
	return c.typ
}

// Gets the component's implementation type.
// Returns string
// the component's implementation type.
func (c *Descriptor) Kind() string {
	return c.kind
}

// Gets the unique component's name.
// Returns string
// the unique component's name.
func (c *Descriptor) Name() string {
	return c.name
}

// Gets the component's implementation version.
// Returns string
// the component's implementation version.
func (c *Descriptor) Version() string {
	return c.version
}

func matchField(field1 string, field2 string) bool {
	return field1 == "" || field2 == "" || field1 == field2
}

// Partially matches this descriptor to another descriptor. Fields that contain "*" or null are excluded from the match.
// see
// exactMatch
// Parameters:
// 			- descriptor *Descriptor
// 			the descriptor to match this one against.
// Returns bool
// true if descriptors match and false otherwise
func (c *Descriptor) Match(descriptor *Descriptor) bool {
	return matchField(c.group, descriptor.Group()) &&
		matchField(c.typ, descriptor.Type()) &&
		matchField(c.kind, descriptor.Kind()) &&
		matchField(c.name, descriptor.Name()) &&
		matchField(c.version, descriptor.Version())
}

func exactMatchField(field1 string, field2 string) bool {
	if field1 == "" && field2 == "" {
		return true
	}
	if field1 == "" || field2 == "" {
		return false
	}
	return field1 == field2
}

// Matches this descriptor to another descriptor by all fields. No exceptions are made.
// see
// Match
// Parameters:
// 			- descriptor *Descriptor
// 			the descriptor to match this one against.
// Returns bool
// true if descriptors match and false otherwise.
func (c *Descriptor) ExactMatch(descriptor *Descriptor) bool {
	return exactMatchField(c.group, descriptor.Group()) &&
		exactMatchField(c.typ, descriptor.Type()) &&
		exactMatchField(c.kind, descriptor.Kind()) &&
		exactMatchField(c.name, descriptor.Name()) &&
		exactMatchField(c.version, descriptor.Version())
}

// Checks whether all descriptor fields are set. If descriptor has at least one "*" or null field it is considered "incomplete",
// Returns bool
// true if all descriptor fields are defined and false otherwise.
func (c *Descriptor) IsComplete() bool {
	return c.group != "" && c.typ != "" && c.kind != "" &&
		c.name != "" && c.version != ""
}

// Compares this descriptor to a value. If value is a Descriptor it tries to match them, otherwise the method returns false.
// see
// Match
// Parameters:
// 			- value interface{}
// 			the value to match against this descriptor.
// Returns bool
// true if the value is matching descriptor and false otherwise.
func (c *Descriptor) Equals(value interface{}) bool {
	descriptor, ok := value.(*Descriptor)
	if ok {
		return c.Match(descriptor)
	}
	return false
}

// Gets a string representation of the object. The result is a colon-separated list of descriptor fields as "mygroup:connector:aws:default:1.0"
// Returns string
// a string representation of the object.
func (c *Descriptor) String() string {
	result := ""
	if c.group == "" {
		result += "*"
	} else {
		result += c.group
	}
	if c.typ == "" {
		result += ":*"
	} else {
		result += ":" + c.typ
	}
	if c.kind == "" {
		result += ":*"
	} else {
		result += ":" + c.kind
	}
	if c.name == "" {
		result += ":*"
	} else {
		result += ":" + c.name
	}
	if c.version == "" {
		result += ":*"
	} else {
		result += ":" + c.version
	}
	return result
}

// Parses colon-separated list of descriptor fields and returns them as a Descriptor.
// throws
// a ConfigError if the descriptor string is of a wrong format.
// Parameters:
// 			- value string
// 			colon-separated descriptor fields to initialize Descriptor.
// Returns *Descriptor
// a newly created Descriptor.
func ParseDescriptorFromString(value string) (*Descriptor, error) {
	if value == "" {
		return nil, nil
	}

	tokens := strings.Split(value, ":")
	if len(tokens) != 5 {
		return nil, errors.NewConfigError("", "BAD_DESCRIPTOR", "Descriptor "+value+" is in wrong format")
	}

	return NewDescriptor(strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1]),
		strings.TrimSpace(tokens[2]), strings.TrimSpace(tokens[3]), strings.TrimSpace(tokens[4])), nil
}
