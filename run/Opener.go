package run

/*
Helper class that opens components.
*/
type TOpener struct{}

var Opener *TOpener = &TOpener{}

// Checks if specified component is opened.
// To be checked components must implement IOpenable interface. If they don't the call to this method returns true.
// see
// IOpenable
//  - component interface{}
//  the component that is to be checked.
// Returns bool
// true if component is opened and false otherwise.
func (c *TOpener) IsOpenOne(component interface{}) bool {
	v, ok := component.(IOpenable)
	if ok {
		return v.IsOpen()
	}
	return true
}

// Checks if all components are opened.
// To be checked components must implement IOpenable interface. If they don't the call to this method returns true.
// see
// isOpenOne
// see
// IOpenable
// Parameters:
// 			- components []interface{}
// 			a list of components that are to be checked.
// Returns bool
// true if all components are opened and false if at least one component is closed.
func (c *TOpener) IsOpen(components []interface{}) bool {
	result := true

	for _, component := range components {
		result = result && c.IsOpenOne(component)
		if !result {
			return result
		}
	}
	return result
}

// Opens specific component.
// To be opened components must implement IOpenable interface. If they don't the call to this method has no effect.
// see
// IOpenable
// Parameters:
// 			- correlationId string
// 			(optional) transaction id to trace execution through call chain.
// 			- component interface{}
// 			the component that is to be opened.
// Returns error
func (c *TOpener) OpenOne(correlationId string, component interface{}) error {
	v, ok := component.(IOpenable)
	if ok {
		return v.Open(correlationId)
	}
	return nil
}

// Opens multiple components.
// To be opened components must implement IOpenable interface. If they don't the call to this method has no effect.
// see
// OpenOne
// see
// IOpenable
// Parameters:
// 			- correlationId string
// 			transaction id to trace execution through call chain.
// 			components []interface{}
// 			the list of components that are to be closed.
// Returns error
func (c *TOpener) Open(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.OpenOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
