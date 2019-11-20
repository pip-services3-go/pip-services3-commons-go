package run

/*
Interface for components that can be asynchronously notified. The notification may include optional argument that describe the occured event.

see
Notifier

see
IExecutable

Example:
type MyComponent {}
    ...
    func (mc *MyComponent)Notify(correlationId: string, args: Parameters){
        fmt.Println("Occured event " + args.GetAsString("event"));
    }

myComponent := MyComponent{};

myComponent.Notify("123", NewParametersFromTuples("event", "Test Event"));
*/
type INotifiable interface {
	// Notifies the component about occured event.
	// Parameters:
	// 			- correlationId string
	// 			transaction id to trace execution through call chain.
	// 			- args *Parameters
	// 			notification arguments.
	Notify(correlationId string, args *Parameters)
}
