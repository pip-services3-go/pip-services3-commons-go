package commands

import (
	"github.com/pip-services-go/pip-services-commons-go/run"
	"github.com/pip-services-go/pip-services-commons-go/validate"
)

type Command struct {
	schema   *validate.Schema
	function func(correlationId string, args *run.Parameters) (interface{}, error)
	name     string
}

func NewCommand(name string, schema *validate.Schema,
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
	// if c.schema != nil {
	// 	err := c.schema.ValidateAndReturnException(correlationId, args)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// Todo: Intercept panic
	result, err := c.function(correlationId, args)

	// let err = new InvocationException(
	// 	correlationId,
	// 	"EXEC_FAILED",
	// 	"Execution " + this.getName() + " failed: " + ex
	// ).withDetails("command", this.getName()).wrap(ex);

	return result, err
}

func (c *Command) Validate(args *run.Parameters) []*validate.ValidationResult {
	// if c.schema != nil {
	// 	return c.schema.Validate(args)
	// }

	return []*validate.ValidationResult{}
}
