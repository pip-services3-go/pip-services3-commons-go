package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/v3/data"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/run"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/validate"
)

/*
Contains a set of commands and events supported by a commandable object.
The CommandSet supports command interceptors to extend and the command call chain.

CommandSets can be used as alternative commandable interface to a business object.
It can be used to auto generate multiple external services for the business object without writing much code.
see
Command
see
Event
see
ICommandable

Example:
type MyDataCommandSet {
	CommandSet
    _controller IMyDataController
}
    func (dcs * MyDataCommandSet) CreateMyDataCommandSet(controller IMyDataController) { // Any data controller interface
        dcs._controller = controller
        dcs.addCommand(dcs.makeGetMyDataCommand())
    }

    func (dcs * MyDataCommandSet) makeGetMyDataCommand() ICommand {
        return NewCommand(
          'get_mydata',
          null,
          (correlationId: string, args: Parameters, func (correlationId string, args *run.Parameters)(interface{}, err) {
              let param = args.getAsString('param');
              return dcs._controller.getMyData(correlationId, param,);
          }
        );
    }
*/
type CommandSet struct {
	commands       []ICommand
	events         []IEvent
	interceptors   []ICommandInterceptor
	commandsByName map[string]ICommand
	eventsByName   map[string]IEvent
}

// Creates an empty CommandSet object.
// Returns *CommandSet
func NewCommandSet() *CommandSet {
	return &CommandSet{
		commands:       []ICommand{},
		events:         []IEvent{},
		interceptors:   []ICommandInterceptor{},
		commandsByName: map[string]ICommand{},
		eventsByName:   map[string]IEvent{},
	}
}

// Gets all commands registered in this command set.
// see
// ICommand
// Returns []ICommand
// a list of commands.
func (c *CommandSet) Commands() []ICommand {
	return c.commands
}

// Gets all events registred in this command set.
// see
// IEvent
// Returns []IEvent
// a list of events.
func (c *CommandSet) Events() []IEvent {
	return c.events
}

// Searches for a command by its name.
// see
// ICommand
// Parameters:
// 			- commandName: string
// 			the name of the command to search for.
// Returns ICommand
// the command, whose name matches the provided name.
func (c *CommandSet) FindCommand(commandName string) ICommand {
	return c.commandsByName[commandName]
}

// Searches for an event by its name in this command set.
// see
// IEvent
// Parameters:
//  		- eventName: string
// 			the name of the event to search for.
// Returns IEvent
// the event, whose name matches the provided name.
func (c *CommandSet) FindEvent(eventName string) IEvent {
	return c.eventsByName[eventName]
}

func (c *CommandSet) buildCommandChain(command ICommand) {
	next := command

	for i := len(c.interceptors) - 1; i >= 0; i-- {
		next = NewInterceptedCommand(c.interceptors[i], next)
	}

	c.commandsByName[next.Name()] = next
}

func (c *CommandSet) rebuildAllCommandChains() {
	c.commandsByName = map[string]ICommand{}

	for _, command := range c.commands {
		c.buildCommandChain(command)
	}
}

// Adds a command to this command set.
// see
// ICommand
// Parameters:
// 			- command: ICommand
// 			 the command to add.
func (c *CommandSet) AddCommand(command ICommand) {
	c.commands = append(c.commands, command)
	c.buildCommandChain(command)
}

// Adds multiple commands to this command set.
// see
// ICommand
// Parameters:
// 			 - commands: []ICommand
// 			the array of commands to add.
func (c *CommandSet) AddCommands(commands []ICommand) {
	for _, command := range commands {
		c.AddCommand(command)
	}
}

// Adds an event to this command set.
// see
// IEvent
// Parameters:
// 				- event: IEvent
// 				the event to add.
func (c *CommandSet) AddEvent(event IEvent) {
	c.events = append(c.events, event)
	c.eventsByName[event.Name()] = event
}

// Adds multiple events to this command set.
// see
// IEvent
// Parameters:
// 			- events: []IEvent
// 			the array of events to add.
func (c *CommandSet) AddEvents(events []IEvent) {
	for _, event := range events {
		c.AddEvent(event)
	}
}

// Adds all of the commands and events from specified command set into this one.
// Parameters:
// 			- commandSet: *CommandSet
// 			the CommandSet to add.
func (c *CommandSet) AddCommandSet(commandSet *CommandSet) {
	c.AddCommands(commandSet.Commands())
	c.AddEvents(commandSet.Events())
}

// Adds a listener to receive notifications on fired events.
// see
// IEventListener
// Parameters:
// 			 - listener: IEventListener
// 				the listener to add.
func (c *CommandSet) AddListener(listener IEventListener) {
	for _, event := range c.events {
		event.AddListener(listener)
	}
}

// Removes previosly added listener.
// see
// IEventListener
// Parameters:
// 			- listener: IEventListener
// 			the listener to remove.
func (c *CommandSet) RemoveListener(listener IEventListener) {
	for _, event := range c.events {
		event.RemoveListener(listener)
	}
}

// Adds a command interceptor to this command set.
// see
// ICommandInterceptor
// Parameters:
// 			-interceptor: ICommandInterceptor
// 			the interceptor to add.
func (c *CommandSet) AddInterceptor(interceptor ICommandInterceptor) {
	c.interceptors = append(c.interceptors, interceptor)
	c.rebuildAllCommandChains()
}

// Executes a command specificed by its name.
// see
// ICommand
// see
// Parameters
// Parameters:
// 			 - correlationId: string
// 			 (optional) transaction id to trace execution through call chain.
// 			 - commandName: string
// 			 the name of that command that is to be executed.
//           - args: Parameters
// 			  the parameters (arguments) to pass to the command for execution.
// return
// result: interface{}
// err: error

func (c *CommandSet) Execute(correlationId string, commandName string, args *run.Parameters) (result interface{}, err error) {
	cref := c.FindCommand(commandName)

	if cref == nil {
		err := errors.NewBadRequestError(
			correlationId,
			"CMD_NOT_FOUND",
			"Request command does not exist",
		).WithDetails("command", commandName)
		return nil, err
	}

	if correlationId == "" {
		correlationId = data.IdGenerator.NextShort()
	}

	// Validate parameters
	results := cref.Validate(args)
	if results != nil && len(results) > 0 {
		err := validate.NewValidationErrorFromResults(correlationId, results, false)
		return nil, err
	}

	return cref.Execute(correlationId, args)
}

// Validates args for command specified by its name using defined schema. If validation schema is
// not defined than the methods returns no errors. It returns validation error if the command is not found.
// see
// Command
// see
// Parameters
// see
// ValidationResult
// Parameters:
// 				- commandName: string
// 				the name of the command for which the 'args' must be validated.
//  			- args: Parameters
// 				the parameters (arguments) to validate.
// Returns []ValidationResult
// an array of ValidationResults. If no command is found by the given name,
// then the returned array of ValidationResults will contain a single entry,
// whose type will be ValidationResultType.Error.
func (c *CommandSet) Validate(commandName string, args *run.Parameters) []*validate.ValidationResult {
	cref := c.FindCommand(commandName)

	if cref == nil {
		return []*validate.ValidationResult{
			validate.NewValidationResult(
				"",
				validate.Error,
				"CMD_NOT_FOUND",
				"Requested command does not exist",
				nil,
				nil,
			),
		}
	}

	return cref.Validate(args)
}

// Fires event specified by its name and notifies all registered listeners
// Parameters:
// 				- correlationId: string
// 				(optional) transaction id to trace execution through call chain.
// 				- eventName: string
// 				the name of the event that is to be fired.
// 				- args: Parameters
// 				the event arguments (parameters).
func (c *CommandSet) Notify(correlationId string, eventName string, args *run.Parameters) {
	event := c.FindEvent(eventName)

	if event != nil {
		event.Notify(correlationId, args)
	}
}
