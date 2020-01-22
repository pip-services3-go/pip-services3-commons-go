package commands

import "github.com/pip-services3-go/pip-services3-commons-go/run"

/*
An interface for listener objects that receive notifications on fired events.
see
IEvent
see
Event

Example:
type MyListener {
   msg string;
}

 func (l* MyListener) onEvent(correlationId string, event IEvent, args Parameters) {
        fmt.Println("Fired event " + event.Name());
    }

let event = NewEvent("myevent");
_listener := MyListener{};
event.addListener(_listener);
event.notify("123", Parameters.FromTuples("param1", "ABC"));

// Console output: Fired event myevent
*/

type IEventListener interface {

	// A method called when events this listener is subscrubed to are fired.
	// Parameters:
	// 			- correlationId: string
	// 			(optional) transaction id to trace execution through call chain.
	// 			- e: IEvent
	// 				a fired evemt
	//  		- value: *run.Parameters
	// 				event arguments.
	OnEvent(correlationId string, e IEvent, value *run.Parameters)
}
