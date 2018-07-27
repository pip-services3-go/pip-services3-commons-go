package run

type IExecutable interface {
	Execute(correlationId string, args *Parameters) (result interface{}, err error)
}
