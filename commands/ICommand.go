package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type ICommand interface {
	run.IExecutable

	Name() string
	Validate(args *run.Parameters) []*validate.ValidationResult
}
