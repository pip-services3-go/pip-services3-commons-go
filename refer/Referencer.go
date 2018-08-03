package refer

type TReferencer struct{}

var Referencer *TReferencer = &TReferencer{}

func (c *TReferencer) SetReferencesForOne(references IReferences, component interface{}) {
	v, ok := component.(IReferenceable)
	if ok {
		v.SetReferences(references)
	}
}

func (c *TReferencer) SetReferences(references IReferences, components []interface{}) {
	for _, component := range components {
		c.SetReferencesForOne(references, component)
	}
}

func (c *TReferencer) UnsetReferencesForOne(component interface{}) {
	v, ok := component.(IUnreferenceable)
	if ok {
		v.UnsetReferences()
	}
}

func (c *TReferencer) UnsetReferences(components []interface{}) {
	for _, component := range components {
		c.UnsetReferencesForOne(component)
	}
}
