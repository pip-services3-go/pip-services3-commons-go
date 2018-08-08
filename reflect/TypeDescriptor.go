package reflect

import (
	"strings"
)

type TypeDescriptor struct {
	name string
	pkg  string
}

func NewTypeDescriptor(name string, pkg string) *TypeDescriptor {
	return &TypeDescriptor{
		name: name,
		pkg:  pkg,
	}
}

func (c *TypeDescriptor) Name() string {
	return c.name
}

func (c *TypeDescriptor) Package() string {
	return c.pkg
}

func (c *TypeDescriptor) Equals(obj interface{}) bool {
	descriptor, ok := obj.(TypeDescriptor)
	if ok {
		if strings.Compare(c.name, descriptor.name) != 0 {
			return false
		}

		if strings.Compare(c.pkg, descriptor.pkg) == 0 {
			return true
		}
	}

	return false
}

func (c *TypeDescriptor) String() string {
	builder := strings.Builder{}

	builder.WriteString(c.name)

	if c.pkg != "" {
		builder.WriteString(",")
		builder.WriteString(c.pkg)
	}

	return builder.String()
}

func ParseTypeDescriptorFromString(value string) (*TypeDescriptor, bool) {
	if value == "" {
		return nil, false
	}

	tokens := strings.Split(value, ",")

	if len(tokens) == 1 {
		return NewTypeDescriptor(strings.Trim(tokens[0], " "), ""), true
	} else if len(tokens) == 2 {
		return NewTypeDescriptor(strings.Trim(tokens[0], " "), strings.Trim(tokens[1], " ")), true
	}

	return nil, false
}
