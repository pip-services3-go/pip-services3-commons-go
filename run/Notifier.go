package run

type TNotifier struct{}

var Notifier *TNotifier = &TNotifier{}

func (c *TNotifier) NotifyOne(correlationId string, component interface{}, args *Parameters) {
	v, ok := component.(INotifiable)
	if ok {
		v.Notify(correlationId, args)
	}
}

func (c *TNotifier) Notify(correlationId string, components []interface{}, args *Parameters) {
	for _, component := range components {
		c.NotifyOne(correlationId, component, args)
	}
}
