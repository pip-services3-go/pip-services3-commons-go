package commands

import (
	"github.com/pip-services-go/pip-services-commons-go/run"
)

type Event struct {
	name      string
	listeners []IEventListener
}

func NewEvent(name string) *Event {
	if name == "" {
		panic("Name cannot be empty")
	}

	return &Event{
		name:      name,
		listeners: []IEventListener{},
	}
}

func (c *Event) Name() string {
	return c.name
}

func (c *Event) Listeners() []IEventListener {
	return c.listeners
}

func (c *Event) AddListener(listener IEventListener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Event) RemoveListener(listener IEventListener) {
	for i, l := range c.listeners {
		if listener == l {
			c.listeners = append(c.listeners[:i], c.listeners[i+1:]...)
			break
		}
	}
}

func (c *Event) Notify(correlationId string, args *run.Parameters) {
	for _, listener := range c.listeners {
		listener.OnEvent(correlationId, c, args)
	}
}
