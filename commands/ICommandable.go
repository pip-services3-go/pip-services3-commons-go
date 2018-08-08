package commands

type ICommandable interface {
	GetCommandSet() *CommandSet
}
