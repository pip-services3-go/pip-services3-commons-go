package run

/*
Interface for components that should clean their state.

Cleaning state most often is used during testing. But there may be situations when it can be done in production.

see
Cleaner

Example:
type MyObjectWithState {
	_state interface{}
}
    ...
    func (mo * MyObjectWithState ) clear(correlationId string) {
        mo._state = interface{}
    }

*/
type ICleanable interface {
	// Clears component state.
	// Parameters:
	// 			- correlationId string
	// 			 transaction id to trace execution through call chain.
	Clear(correlationId string) error
}
