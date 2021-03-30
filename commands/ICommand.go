package commands

import (
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-commons-go/validate"
)

/*
An interface for Commands, which are part of the Command design pattern.
Each command wraps a method or function and allows to call them in uniform and safe manner.
*/

type ICommand interface {
	run.IExecutable
	// Gets the command name.
	// Returns string
	// the command name.
	Name() string
	// Validates command arguments before execution using defined schema.
	// see
	// Parameters
	// see
	// ValidationResult
	// Parameters:
	//  - args: Parameters
	//  the parameters (arguments) to validate.
	// Returns ValidationResult[]
	// an array of ValidationResults.
	Validate(args *run.Parameters) []*validate.ValidationResult
}
