package run

/*
Interface for components that require explicit closure.

For components that require opening as well as closing use IOpenable interface instead.

see
IOpenable

see
Closer

Example:
type MyConnector {
    _client interface{}
}
    ... // The _client can be lazy created

    func (mc *MyConnector) Close(correlationId: string):error {
        if (mc._client != nil) {
            mc._client.Close()
			mc._client = nil
			return nil
		}
    }

*/
type IClosable interface {
	Close(correlationId string) error
}
