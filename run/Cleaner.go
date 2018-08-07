package run

type TCleaner struct{}

var Cleaner *TCleaner = &TCleaner{}

func (c *TCleaner) ClearOne(correlationId string, component interface{}) error {
	v, ok := component.(ICleanable)
	if ok {
		return v.Clear(correlationId)
	}
	return nil
}

func (c *TCleaner) Clear(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.ClearOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
