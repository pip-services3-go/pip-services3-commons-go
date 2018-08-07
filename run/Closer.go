package run

type TCloser struct{}

var Closer *TCloser = &TCloser{}

func (c *TCloser) CloseOne(correlationId string, component interface{}) error {
	v, ok := component.(IClosable)
	if ok {
		return v.Close(correlationId)
	}
	return nil
}

func (c *TCloser) Close(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.CloseOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
