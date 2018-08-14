package validate

import (
	"strings"

	"github.com/pip-services-go/pip-services-commons-go/errors"
)

func NewValidationError(correlationId string, message string, results []*ValidationResult) *errors.ApplicationError {
	if message == "" {
		message = composeErrorMessage(results)
	}
	e := errors.NewBadRequestError(correlationId, "INVALID_DATA", message)
	if results != nil && len(results) > 0 {
		e.WithDetails("results", results)
	}
	return e
}

func composeErrorMessage(results []*ValidationResult) string {
	builder := strings.Builder{}
	builder.WriteString("Validation failed")

	if results == nil || len(results) == 0 {
		return builder.String()
	}

	first := true
	for _, result := range results {
		if result.Type() == Information {
			continue
		}

		if first {
			builder.WriteString(": ")
		} else {
			builder.WriteString(", ")
		}
		builder.WriteString(result.Message())
		first = false
	}

	return builder.String()
}

func NewValidationErrorFromResults(correlationId string, results []*ValidationResult, strict bool) *errors.ApplicationError {
	hasErrors := false

	for _, result := range results {
		if result.Type() == Error {
			hasErrors = true
		}

		if strict && result.Type() == Warning {
			hasErrors = true
		}
	}

	if hasErrors {
		return NewValidationError(correlationId, "", results)
	}

	return nil
}

func ThrowValidationErrorIfNeeded(correlationId string, results []*ValidationResult, strict bool) {
	err := NewValidationErrorFromResults(correlationId, results, strict)
	if err != nil {
		panic(err)
	}
}
