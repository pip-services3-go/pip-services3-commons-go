package run

type IOpenable interface {
	IClosable

	IsOpen() bool
	Open(correlationId string) error
}
