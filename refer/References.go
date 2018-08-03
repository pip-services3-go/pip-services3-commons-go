package refer

type References struct {
	references []*Reference
}

func NewEmptyReferences() *References {
	return &References{
		references: make([]*Reference, 0, 10),
	}
}

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

func (c *References) Put(locator interface{}, component interface{}) {
	if component == nil {
		panic("Component cannot be null")
	}

	reference := NewReference(locator, component)
	c.references = append(c.references, reference)
}

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

func (c *References) GetAllLocators() []interface{} {
	components := make([]interface{}, len(c.references), len(c.references))

	for index, reference := range c.references {
		components[index] = reference.Locator()
	}

	return components
}

func (c *References) GetAll() []interface{} {
	components := make([]interface{}, len(c.references), len(c.references))

	for index, reference := range c.references {
		components[index] = reference.Component()
	}

	return components
}

func (c *References) GetOneOptional(locator interface{}) interface{} {
	components, err := c.Find(locator, true)
	if err != nil || len(components) == 0 {
		return nil
	}
	return components[0]
}

func (c *References) GetOneRequired(locator interface{}) (interface{}, error) {
	components, err := c.Find(locator, true)
	if err != nil || len(components) == 0 {
		return nil, err
	}
	return components[0], nil
}

func (c *References) GetOptional(locator interface{}) []interface{} {
	components, _ := c.Find(locator, false)
	return components
}

func (c *References) GetRequired(locator interface{}) ([]interface{}, error) {
	return c.Find(locator, true)
}

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

func NewReferencesFromTuples(tuples ...interface{}) *References {
	return NewReferences(tuples)
}
