package refer

/*
Error when required component dependency cannot be found.
*/
import (
	"fmt"

	"github.com/pip-services3-go/pip-services3-commons-go/v3/errors"
)

// Creates an error instance and assigns its values.
// Parameters:
// 			- correlationId string
// 			- locator interface{}
// 			the locator to find reference to dependent component.
// Returns *errors.ApplicationError
func NewReferenceError(correlationId string, locator interface{}) *errors.ApplicationError {
	message := fmt.Sprintf("Failed to obtain reference to %v", locator)
	e := errors.NewInternalError(correlationId, "REF_ERROR", message)
	e.WithDetails("locator", locator)
	return e
}
