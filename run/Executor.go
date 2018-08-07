package run

type TExecutor struct{}

var Executor *TExecutor = &TExecutor{}

func (c *TExecutor) ExecuteOne(correlationId string, component interface{}, args *Parameters) (interface{}, error) {
	v, ok := component.(IExecutable)
	if ok {
		return v.Execute(correlationId, args)
	}
	return nil, nil
}

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
