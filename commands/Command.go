package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

/*
Concrete implementation of ICommand interface. Command allows to call a
method or function using Command pattern.
Example:

	command := NewCommand("add", null, func (correlationId string, args *run.Parameters)(interface{}, err) {
		param1 := args.getAsFloat("param1");
    	param2 := args.getAsFloat("param2");
    	return (param1 + param2), nil;
	});

	result, err := command.Execute("123", Parameters.NewParametersFromTuples("param1", 2, "param2", 2))
	if (err) {
		fmt.Println(err)
	} else {
		fmt.Println("2 + 2 = " + result)
	}
// Console output: 2 + 2 = 4
//
*/
type Command struct {
	schema   validate.ISchema
	function func(correlationId string, args *run.Parameters) (interface{}, error)
	name     string
}

//Creates a new command object and assigns it's parameters.
//
//Parameters
//				- name: string - the command name.
// 				- schema: validate.ISchema the schema to validate command arguments.
//				function: func(correlationId string, args *run.Parameters) (interface{}, error)
//				the function to be executed by this command.
//
//Returns *Command
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

//Gets the command name.
//
//Returns string - the name of this command.
func (c *Command) Name() string {
	return c.name
}

//Executes the command. Before execution it validates args using the defined schema.
//The command execution intercepts exceptions raised by the called function and returns them as an error
//in callback.
//Parameters:
//correlationId: string - (optional) transaction id to trace execution through call chain.
//
//args: run.Parameters - the parameters (arguments) to pass to this command for execution.
//
//Returns (interface{}, error)
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

//
//Validates the command args before execution using the defined schema.
//
//Parameters:
//args: run.Parameters - the parameters (arguments) to validate using this command's schema.
//Returns []*validate.ValidationResult an array of ValidationResults or an empty array (if no schema is set).
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

// GetSchema methods return validation schema for this command
func (c *Command) GetSchema() validate.ISchema {
	return c.schema
}
