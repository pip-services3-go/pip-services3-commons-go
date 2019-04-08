package commands

import "github.com/pip-services3-go/pip-services3-commons-go/run"

type IEvent interface {
	run.INotifiable

	Name() string
	Listeners() []IEventListener
	AddListener(listener IEventListener)
	RemoveListener(listener IEventListener)
}
