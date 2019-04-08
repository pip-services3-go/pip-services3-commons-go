package refer

import (
	"fmt"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

func NewReferenceError(correlationId string, locator interface{}) *errors.ApplicationError {
	message := fmt.Sprintf("Failed to obtain reference to %v", locator)
	e := errors.NewInternalError(correlationId, "REF_ERROR", message)
	e.WithDetails("locator", locator)
	return e
}
