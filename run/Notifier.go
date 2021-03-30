package run

/*
Helper class that notifies components.
*/
type TNotifier struct{}

var Notifier *TNotifier = &TNotifier{}

// Notifies specific component.
// To be notiied components must implement INotifiable interface. If they don't the call to this method has no effect.
// see
// INotifiable
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - component interface{}
//  the component that is to be notified.
//  - args *Parameters
//  notifiation arguments.
func (c *TNotifier) NotifyOne(correlationId string, component interface{}, args *Parameters) {
	v, ok := component.(INotifiable)
	if ok {
		v.Notify(correlationId, args)
	}
}

// Notifies multiple components.
// To be notified components must implement INotifiable interface. If they don't the call to this method has no effect.
// see
// NotifyOne
// see
// INotifiable
// Parameters:
// 			- correlationId string
// 			 transaction id to trace execution through call chain.
// 			- components []interface{}
// 			a list of components that are to be notified.
// 			- args *Parameters
// 			notification arguments.
func (c *TNotifier) Notify(correlationId string, components []interface{}, args *Parameters) {
	for _, component := range components {
		c.NotifyOne(correlationId, component, args)
	}
}
