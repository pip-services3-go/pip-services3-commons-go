package run

type IClosable interface {
	Close(correlationId string) error
}
