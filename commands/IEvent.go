package commands

import "github.com/pip-services-go/pip-services-commons-go/run"

type IEvent interface {
	run.INotifiable

	Name() string
	Listeners() []IEventListener
	AddListener(listener IEventListener)
	RemoveListener(listener IEventListener)
}
