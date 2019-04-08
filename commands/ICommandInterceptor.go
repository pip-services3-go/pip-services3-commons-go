package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type ICommandInterceptor interface {
	Name(command ICommand) string
	Execute(correlationId string, command ICommand, args *run.Parameters) (interface{}, error)
	Validate(command ICommand, args *run.Parameters) []*validate.ValidationResult
}
