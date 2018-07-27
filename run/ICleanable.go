package run

type ICleanable interface {
	Clear(correlationId string) error
}
