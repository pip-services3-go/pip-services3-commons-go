package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type CommandSet struct {
	commands       []ICommand
	events         []IEvent
	interceptors   []ICommandInterceptor
	commandsByName map[string]ICommand
	eventsByName   map[string]IEvent
}

func NewCommandSet() *CommandSet {
	return &CommandSet{
		commands:       []ICommand{},
		events:         []IEvent{},
		interceptors:   []ICommandInterceptor{},
		commandsByName: map[string]ICommand{},
		eventsByName:   map[string]IEvent{},
	}
}

func (c *CommandSet) Commands() []ICommand {
	return c.commands
}

func (c *CommandSet) Events() []IEvent {
	return c.events
}

func (c *CommandSet) FindCommand(commandName string) ICommand {
	return c.commandsByName[commandName]
}

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

func (c *CommandSet) AddCommand(command ICommand) {
	c.commands = append(c.commands, command)
	c.buildCommandChain(command)
}

func (c *CommandSet) AddCommands(commands []ICommand) {
	for _, command := range commands {
		c.AddCommand(command)
	}
}

func (c *CommandSet) AddEvent(event IEvent) {
	c.events = append(c.events, event)
	c.eventsByName[event.Name()] = event
}

func (c *CommandSet) AddEvents(events []IEvent) {
	for _, event := range events {
		c.AddEvent(event)
	}
}

func (c *CommandSet) AddCommandSet(commandSet *CommandSet) {
	c.AddCommands(commandSet.Commands())
	c.AddEvents(commandSet.Events())
}

func (c *CommandSet) AddListener(listener IEventListener) {
	for _, event := range c.events {
		event.AddListener(listener)
	}
}

func (c *CommandSet) RemoveListener(listener IEventListener) {
	for _, event := range c.events {
		event.RemoveListener(listener)
	}
}

func (c *CommandSet) AddInterceptor(interceptor ICommandInterceptor) {
	c.interceptors = append(c.interceptors, interceptor)
	c.rebuildAllCommandChains()
}

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

func (c *CommandSet) Notify(correlationId string, eventName string, args *run.Parameters) {
	event := c.FindEvent(eventName)

	if event != nil {
		event.Notify(correlationId, args)
	}
}
