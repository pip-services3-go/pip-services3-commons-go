package run

/*
Interface for components that require explicit opening and closing.

For components that perform opening on demand consider using [[ICloseable]] interface instead.

see
IOpenable

see
Opener

Example:
type MyPersistence {
	_client interface{}
}

    func (mp* MyPersistence)IsOpen() bool {
        return mp._client != nil;
    }

    (mp* MyPersistence) Open(correlationId: string) error {
        if (mp.isOpen()) {
            return nil;
        }
    }

    (mp* MyPersistence) Close(correlationId: string) {
        if (mp._client != nil) {
            mp._client.close();
            mp._client = nil;
        }
    }

*/
type IOpenable interface {
	IClosable
	// Checks if the component is opened.
	// Returns bool
	// true if the component has been opened and false otherwise.
	IsOpen() bool
	// Opens the component.
	// Parameters:
	// 			- correlationId: string
	// 			transaction id to trace execution through call chain.
	// Return error
	Open(correlationId string) error
}
