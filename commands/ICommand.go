package commands

import (
	"github.com/pip-services-go/pip-services-commons-go/run"
	"github.com/pip-services-go/pip-services-commons-go/validate"
)

type ICommand interface {
	run.IExecutable

	Name() string
	Validate(args *run.Parameters) []*validate.ValidationResult
}
