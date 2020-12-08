package refer

/*
The most basic implementation of IReferences to store and locate component references.

see
IReferences

Example:
 type MyController  {
 	_persistence IMyPersistence;
 }

 func (mc *MyController) setReferences(references IReferences) {
     mc._persistence = references.GetOneRequired(
         NewDescriptor("mygroup", "persistence", "*", "*", "1.0")
     );
 }

 persistence := NewMyMongoDbPersistence();
 
 controller := MyController();
 
 references := NewReferencesFromTuples(
     new Descriptor("mygroup", "persistence", "mongodb", "default", "1.0"), persistence,
     new Descriptor("mygroup", "controller", "default", "default", "1.0"), controller
 );
 controller.setReferences(references);
*/
type References struct {
	references []*Reference
}

// Creates a new instance of references and initializes it with references.
// Returns *References
func NewEmptyReferences() *References {
	return &References{
		references: make([]*Reference, 0, 10),
	}
}

// Creates a new instance of references and initializes it with references.
// Parameters:
//  - tuples []interface{}
//  a list of values where odd elements are locators and the following even elements are component references
// Returns *References
func NewReferences(tuples []interface{}) *References {
	c := NewEmptyReferences()

	for index := 0; index < len(tuples); index += 2 {
		if index+1 >= len(tuples) {
			break
		}

		c.Put(tuples[index], tuples[index+1])
	}

	return c
}

// Puts a new reference into this reference map.
// Parameters:
//  - locator interface{}
//  a locator to find the reference by.
//  - component interface{}
//  a component reference to be added.
func (c *References) Put(locator interface{}, component interface{}) {
	if component == nil {
		panic("Component cannot be null")
	}

	reference := NewReference(locator, component)
	c.references = append(c.references, reference)
}

// Removes a previously added reference that matches specified locator. If many references match the locator, it removes only the first one. When all references shall be removed, use removeAll method instead.
// see
// RemoveAll
// Parameters:
//  - locator interface{}
//  a locator to remove reference
// Returns interface{}
// the removed component reference.
func (c *References) Remove(locator interface{}) interface{} {
	if locator == nil {
		return nil
	}

	for index := len(c.references) - 1; index >= 0; index-- {
		reference := c.references[index]
		if reference.Match(locator) {
			c.references = append(c.references[:index], c.references[index+1:]...)
			return reference.Component()
		}
	}

	return nil
}

// Removes all component references that match the specified locator.
// Parameters:
//  - locator interface{}
//  the locator to remove references by.
// Returns []interface{}
// a list, containing all removed references.
func (c *References) RemoveAll(locator interface{}) []interface{} {
	components := make([]interface{}, 0, 5)

	if locator == nil {
		return components
	}

	for index := len(c.references) - 1; index >= 0; index-- {
		reference := c.references[index]
		if reference.Match(locator) {
			c.references = append(c.references[:index], c.references[index+1:]...)
			components = append(components, reference.Component())
		}
	}

	return components
}

// Gets locators for all registered component references in this reference map.
// Returns []interface{}
// a list with component locators.
func (c *References) GetAllLocators() []interface{} {
	components := make([]interface{}, len(c.references), len(c.references))

	for index, reference := range c.references {
		components[index] = reference.Locator()
	}

	return components
}

// Gets all component references registered in this reference map.
// Returns []interface{}
// a list with component references.
func (c *References) GetAll() []interface{} {
	components := make([]interface{}, len(c.references), len(c.references))

	for index, reference := range c.references {
		components[index] = reference.Component()
	}

	return components
}

// Gets an optional component reference that matches specified locator.
// Parameters:
//  - locator interface{}
//  the locator to find references by.
// Returns interface{}
// a matching component reference or nil if nothing was found.
func (c *References) GetOneOptional(locator interface{}) interface{} {
	components, err := c.Find(locator, false)
	if err != nil || len(components) == 0 {
		return nil
	}
	return components[0]
}

// Gets a required component reference that matches specified locator.
// throws
// a ReferenceError when no references found.
// Parameters:
//  - locator interface{}
//  the locator to find a reference by.
// Returns interface{}
// a matching component reference.
func (c *References) GetOneRequired(locator interface{}) (interface{}, error) {
	components, err := c.Find(locator, true)
	if err != nil || len(components) == 0 {
		return nil, err
	}
	return components[0], nil
}

// Gets all component references that match specified locator.
// Parameters:
//  - locator interface{}
//  the locator to find references by.
// Returns []interface{}
// a list with matching component references or empty list if nothing was found.
func (c *References) GetOptional(locator interface{}) []interface{} {
	components, _ := c.Find(locator, false)
	return components
}

// Gets all component references that match specified locator. At least one component reference must be present. If it doesn't the method throws an error.
// throws
// a ReferenceError when no references found.
// Parameters:
//  - locator interface{}
//  the locator to find references by.
// Returns []interface{}
// a list with matching component references.
func (c *References) GetRequired(locator interface{}) ([]interface{}, error) {
	return c.Find(locator, true)
}

// Gets all component references that match specified locator.
// throws
// a ReferenceError when required is set to true but no references found.
// Parameters:
//  - locator interface{}
//  the locator to find a reference by.
//  - required bool
//  forces to raise an exception if no reference is found.
// Returns []interface{}
// a list with matching component references.
func (c *References) Find(locator interface{}, required bool) ([]interface{}, error) {
	if locator == nil {
		panic("Locator cannot be null")
	}

	components := make([]interface{}, 0, 2)

	// Search all references
	for index := len(c.references) - 1; index >= 0; index-- {
		reference := c.references[index]
		if reference.Match(locator) {
			component := reference.Component()
			components = append(components, component)
		}
	}

	if len(components) == 0 && required {
		err := NewReferenceError("", locator)
		return components, err
	}

	return components, nil
}

// Creates a new References from a list of key-value pairs called tuples.
// Parameters:
//  - tuples  ...interface{}
//  a list of values where odd elements are locators and the following even elements are component references
// Returns *References
// a newly created References.
func NewReferencesFromTuples(tuples ...interface{}) *References {
	return NewReferences(tuples)
}
