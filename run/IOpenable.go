package run

type IOpenable interface {
	IClosable

	IsOpened() bool
	Open(correlationId string) error
}
