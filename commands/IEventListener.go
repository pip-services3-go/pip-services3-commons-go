package commands

import "github.com/pip-services3-go/pip-services3-commons-go/run"

type IEventListener interface {
	OnEvent(correlationId string, e IEvent, value *run.Parameters)
}
