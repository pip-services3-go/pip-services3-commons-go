package run

/*
Interface for components that require execution parameters.
*/
type IParameterized interface {
	// Sets execution parameters.
	// Parameters:
	// parameters *Parameters
	// execution parameters.
	SetParameters(parameters *Parameters)
}
