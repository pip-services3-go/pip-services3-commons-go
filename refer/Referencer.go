package refer

/*
Helper class that sets and unsets references to components.
*/
type TReferencer struct{}

var Referencer *TReferencer = &TReferencer{}

// Sets references to specific component.
// To set references components must implement IReferenceable interface. If they don't the call to this method has no effect.
// see
// IReferenceable
// Parameters:
//  - references IReferences
//  the references to be set.
//  - component interface{}
//  the component to set references to.
func (c *TReferencer) SetReferencesForOne(references IReferences, component interface{}) {
	v, ok := component.(IReferenceable)
	if ok {
		v.SetReferences(references)
	}
}

// Sets references to multiple components.
// To set references components must implement IReferenceable interface. If they don't the call to this method has no effect.
// see
// IReferenceable
// Parameters:
// 			- references IReferences
// 			the references to be set.
// 			- components []interface{}
// 			a list of components to set the references to.
func (c *TReferencer) SetReferences(references IReferences, components []interface{}) {
	for _, component := range components {
		c.SetReferencesForOne(references, component)
	}
}

// Unsets references in specific component.
// To unset references components must implement IUnreferenceable interface. If they don't the call to this method has no effect.
// see
// IUnreferenceable
// Parameters:
// 			 - component interface{}
//           the component to unset references.
func (c *TReferencer) UnsetReferencesForOne(component interface{}) {
	v, ok := component.(IUnreferenceable)
	if ok {
		v.UnsetReferences()
	}
}

// Unsets references in multiple components.
// To unset references components must implement IUnreferenceable interface. If they don't the call to this method has no effect.
// see
// IUnreferenceable
// Parameters:
// 			 - components [] interface{}
// 			 the list of components, whose references must be cleared.
func (c *TReferencer) UnsetReferences(components []interface{}) {
	for _, component := range components {
		c.UnsetReferencesForOne(component)
	}
}
