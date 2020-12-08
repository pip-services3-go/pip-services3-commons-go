package refer

import "github.com/pip-services3-go/pip-services3-commons-go/data"

/*
Contains a reference to a component and locator to find it. It is used by References to store registered component references.
*/
type Reference struct {
	locator   interface{}
	component interface{}
}

// Create a new instance of the reference object and assigns its values.
// Parameters:
//  - locator interface{}
//  a locator to find the reference.
//  - component interface {}
// Returns *Reference
func NewReference(locator interface{}, component interface{}) *Reference {
	if component == nil {
		panic("Component cannot be null")
	}

	return &Reference{
		locator:   locator,
		component: component,
	}
}

// Gets the stored component reference.
// Returns any
// the component's references.
func (c *Reference) Component() interface{} {
	return c.component
}

// Gets the stored component locator.
// Returns any
// the component's locator.
func (c *Reference) Locator() interface{} {
	return c.locator
}

// Matches locator to this reference locator.
// Descriptors are matched using equal method. All other locator types are matched using direct comparison.
// see
// Descriptor
// Parameters:
//  - locator interface{}
//  the locator to match.
// Returns bool
// true if locators are matching and false it they don't.
func (c *Reference) Match(locator interface{}) bool {
	// Check for nil locator
	if locator == nil {
		return false
	}

	// Locate by direct reference matching
	if c.component == locator {
		return true
	}

	// Locate by direct locator matching
	equatable, ok := c.locator.(data.IEquatable)
	if ok {
		return equatable.Equals(locator)
	}

	// Locate by direct locator matching
	return c.locator == locator
}
