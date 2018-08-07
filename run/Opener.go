package run

type TOpener struct{}

var Opener *TOpener = &TOpener{}

func (c *TOpener) IsOpenOne(component interface{}) bool {
	v, ok := component.(IOpenable)
	if ok {
		return v.IsOpen()
	}
	return true
}

func (c *TOpener) IsOpen(components []interface{}) bool {
	result := true

	for _, component := range components {
		result = result && c.IsOpenOne(component)
		if !result {
			return result
		}
	}
	return result
}

func (c *TOpener) OpenOne(correlationId string, component interface{}) error {
	v, ok := component.(IOpenable)
	if ok {
		return v.Open(correlationId)
	}
	return nil
}

func (c *TOpener) Open(correlationId string, components []interface{}) error {
	for _, component := range components {
		err := c.OpenOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
