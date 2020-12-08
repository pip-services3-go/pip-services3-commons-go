package run

/*
Helper class that cleans stored object state.
*/
type TCleaner struct{}

var Cleaner *TCleaner = &TCleaner{}

// Clears state of specific component.
// To be cleaned state components must implement ICleanable interface. If they don't the call to this method has no effect.
// Parameters:
//  - correlationId: string
//  transaction id to trace execution through call chain.
//  - component interface{}
//  the component that is to be cleaned.
// Returns error
func (c *TCleaner) ClearOne(correlationId string, component interface{}) error {
	v, ok := component.(ICleanable)
	if ok {
		return v.Clear(correlationId)
	}
	return nil
}

// Clears state of multiple components.
// To be cleaned state components must implement ICleanable interface. If they don't the call to this method has no effect.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - components []interface{}
// the list of components that are to be cleaned.
// Returns error
func (c *TCleaner) Clear(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.ClearOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
