package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type Command struct {
	schema   validate.ISchema
	function func(correlationId string, args *run.Parameters) (interface{}, error)
	name     string
}

func NewCommand(name string, schema validate.ISchema,
	function func(correlationId string, args *run.Parameters) (interface{}, error)) *Command {
	if name == "" {
		panic("Name cannot be empty")
	}
	if function == nil {
		panic("Function cannot be nil")
	}

	return &Command{
		name:     name,
		schema:   schema,
		function: function,
	}
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Execute(correlationId string, args *run.Parameters) (interface{}, error) {
	if c.schema != nil {
		err := c.schema.ValidateAndReturnError(correlationId, args, false)
		if err != nil {
			return nil, err
		}
	}

	var err error

	// Execute in inner function to capture errors
	result, err2 := func() (interface{}, error) {
		// Intercepting unhandled errors
		defer func() {
			if r := recover(); r != nil {
				tempMessage := convert.StringConverter.ToString(r)
				tempError := errors.NewInvocationError(
					correlationId,
					"EXEC_FAILED",
					"Execution "+c.Name()+" failed: "+tempMessage,
				).WithDetails("command", c.Name())

				cause, ok := r.(error)
				if ok {
					tempError.WithCause(cause)
				}

				err = tempError
			}
		}()

		return c.function(correlationId, args)
	}()

	if err2 != nil {
		err = err2
	}

	return result, err
}

func (c *Command) Validate(args *run.Parameters) []*validate.ValidationResult {
	if c.schema != nil {
		results := c.schema.Validate(args)
		if results == nil {
			results = []*validate.ValidationResult{}
		}
		return results
	}

	return []*validate.ValidationResult{}
}
