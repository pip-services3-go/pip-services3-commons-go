package run

/*
Helper class that closes previously opened components.
*/
type TCloser struct{}

var Closer *TCloser = &TCloser{}

// Closes specific component.

// To be closed components must implement ICloseable interface. If they don't the call to this method has no effect.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - component interface{}
// 	the component that is to be closed.
// Returns error
func (c *TCloser) CloseOne(correlationId string, component interface{}) error {
	v, ok := component.(IClosable)
	if ok {
		return v.Close(correlationId)
	}
	return nil
}

// Closes multiple components.
// To be closed components must implement ICloseable interface. If they don't the call to this method has no effect.
// Parameters:
// 			- correlationId string
//  		transaction id to trace execution through call chain.
// 			- components []interface{}
// 			the list of components that are to be closed.
// Returns error
func (c *TCloser) Close(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.CloseOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
