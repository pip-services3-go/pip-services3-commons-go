package run

/*
 Helper class that executes components.
*/
type TExecutor struct{}

var Executor *TExecutor = &TExecutor{}

// Executes specific component.
// To be executed components must implement IExecutable interface. If they don't the call to this method has no effect.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - component interface{}
//  the component that is to be executed.
//  - args: *Parameters
//  execution arguments.
// Returns []interface{}, error
// execution result or error
func (c *TExecutor) ExecuteOne(correlationId string, component interface{}, args *Parameters) (interface{}, error) {
	v, ok := component.(IExecutable)
	if ok {
		return v.Execute(correlationId, args)
	}
	return nil, nil
}

// Executes multiple components.

// To be executed components must implement IExecutable interface. If they don't the call to this method has no effect.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - components []interface{}
//  a list of components that are to be executed.
//  - args *Parameters
//  execution arguments.
// Returns []interface{}, error
// execution result or error
func (c *TExecutor) Execute(correlationId string, components []interface{}, args *Parameters) ([]interface{}, error) {
	results := make([]interface{}, 0, 5)

	for _, component := range components {
		result, err := c.ExecuteOne(correlationId, component, args)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}
