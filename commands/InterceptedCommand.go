package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

/*
Implements a command wrapped by an interceptor. It allows to build command call chains.
The interceptor can alter execution and delegate calls to a next command, which can be intercepted or concrete.
see
ICommand
see
ICommandInterceptor

Example:
 type CommandLogger {
 	msg string
 }

 func (cl * CommandLogger) Name(command ICommand) string {
     return command.Name();
 }

 func (cl * CommandLogger) Execute(correlationId string, command ICommand, args Parameters) (res interface{}, err error){
     fmt.Println("Executed command " + command.Name());
     return command.Execute(correlationId, args);
 }

 func (cl * CommandLogger) Validate(command: ICommand, args: Parameters): ValidationResult[] {
     return command.Validate(args);
 }

 logger := CommandLogger{mgs:"CommandLoger"};
 loggedCommand = NewInterceptedCommand(logger, command);
 
 // Each called command will output: Executed command <command name>
*/
type InterceptedCommand struct {
	interceptor ICommandInterceptor
	next        ICommand
}

// Creates a new InterceptedCommand, which serves as a link in an execution chain.
// Contains information about the interceptor that is being used and the next command in the chain.
// Parameters:
//  - interceptor: ICommandInterceptor
//  the interceptor that is intercepting the command.
//  - next: ICommand
//  (link to) the next command in the command's execution chain.
// Returns *InterceptedCommand
func NewInterceptedCommand(interceptor ICommandInterceptor, next ICommand) *InterceptedCommand {
	return &InterceptedCommand{
		interceptor: interceptor,
		next:        next,
	}
}

// Returns string
// the name of the command that is being intercepted.
func (c *InterceptedCommand) Name() string {
	return c.interceptor.Name(c.next)
}

// Executes the next command in the execution chain using the given parameters (arguments).
// see
// Parameters
// Parameters:
//  - correlationId: string
//  	unique transaction id to trace calls across components.
//  - args: Parameters
//  	the parameters (arguments) to pass to the command for execution.
// Returns:
// err: error
// result: interface{}
func (c *InterceptedCommand) Execute(correlationId string, args *run.Parameters) (result interface{}, err error) {
	return c.interceptor.Execute(correlationId, c.next, args)
}

// Validates the parameters (arguments) that are to be passed to the command that is next in the execution chain.
// see
// Parameters
// see
// ValidationResult
// Parameters:
//  - args: Parameters
//  the parameters (arguments) to validate for the next command.
// Returns []*ValidationResult
// an array of *ValidationResults.
func (c *InterceptedCommand) Validate(args *run.Parameters) []*validate.ValidationResult {
	return c.interceptor.Validate(c.next, args)
}
