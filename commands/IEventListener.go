package commands

import "github.com/pip-services-go/pip-services-commons-go/run"

type IEventListener interface {
	OnEvent(correlationId string, e IEvent, value *run.Parameters)
}
