package run

type INotifiable interface {
	Notify(correlationId string, args *Parameters)
}
