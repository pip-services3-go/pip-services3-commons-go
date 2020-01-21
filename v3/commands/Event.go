package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/v3/run"
)

/*
Concrete implementation of IEvent interface. It allows to send asynchronous
notifications to multiple subscribed listeners.
Example:
event: = NewEvent("my_event");

event.AddListener(myListener);

event.Notify("123", Parameters.fromTuples(
  "param1", "ABC",
  "param2", 123
));
*/

type Event struct {
	name      string
	listeners []IEventListener
}

// Creates a new event and assigns its name.
// throws
// an Error if the name is null.
// Parameters:
// 				- name: string
// 					the name of the event that is to be created.
// Returns Event

func NewEvent(name string) *Event {
	if name == "" {
		panic("Name cannot be empty")
	}

	return &Event{
		name:      name,
		listeners: []IEventListener{},
	}
}

// Gets the name of the event.
// Returns string
// the name of this event.

func (c *Event) Name() string {
	return c.name
}

// Gets all listeners registred in this event.
// Returns []IEventListener
// a list of listeners.

func (c *Event) Listeners() []IEventListener {
	return c.listeners
}

// Adds a listener to receive notifications when this event is fired.
// Parameters:
// 				- listener: IEventListener
// 					the listener reference to add.

func (c *Event) AddListener(listener IEventListener) {
	c.listeners = append(c.listeners, listener)
}

// Removes a listener, so that it no longer receives notifications for this event.
// Parameters:
//  			- listener: IEventListener
// 				the listener reference to remove.

func (c *Event) RemoveListener(listener IEventListener) {
	for i, l := range c.listeners {
		if listener == l {
			c.listeners = append(c.listeners[:i], c.listeners[i+1:]...)
			break
		}
	}
}

// Fires this event and notifies all registred listeners.
// Parameters:
// 				- correlationId: string
// 				(optional) transaction id to trace execution through call chain.
// 				- args: Parameters
// 				the parameters to raise this event with.

func (c *Event) Notify(correlationId string, args *run.Parameters) {
	for _, listener := range c.listeners {
		listener.OnEvent(correlationId, c, args)
	}
}
