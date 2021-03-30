package commands

import "github.com/pip-services3-go/pip-services3-commons-go/run"

/*
An interface for Events, which are part of the Command design pattern.
Events allows to send asynchronious notifications to multiple subscribed listeners.

see
IEventListener
*/

type IEvent interface {
	run.INotifiable

	// Gets the event name.
	// Returns string
	// the name of the event.
	Name() string

	// Gets all subscribed listeners.
	// Returns []IEventListener
	// a list of listeners.
	Listeners() []IEventListener

	// Adds a listener to receive notifications for this event.
	// Parameters:
	//  - listener: IEventListener
	//  the listener reference to add.
	AddListener(listener IEventListener)

	// Removes a listener, so that it no longer receives notifications for this event.
	// Parameters
	//  - listener: IEventListener
	//  the listener reference to remove.
	RemoveListener(listener IEventListener)
}
