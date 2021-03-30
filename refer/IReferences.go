package refer

/*
Interface for a map that holds component references and passes them to components to establish dependencies with each other.

Together with IReferenceable and IUnreferenceable interfaces it implements a Locator pattern that is used by PipServices toolkit for Inversion of Control to assign external dependencies to components.

The IReferences object is a simple map, where keys are locators and values are component references. It allows to add, remove and find components by their locators. Locators can be any values like integers, strings or component types. But most often PipServices toolkit uses Descriptor as locators that match by 5 fields: group, type, kind, name and version.
*/

type IReferences interface {
	// Puts a new reference into this reference map.
	// Parameters:
	//  - locator interface{}
	//  a locator to find the reference by.
	//  - component interface{}
	//  a component reference to be added.
	// Returns interface{}
	Put(locator interface{}, component interface{})
	// 	Removes a previously added reference that matches specified locator. If many references match the locator, it removes only the first one. When all references shall be removed, use removeAll method instead.
	// see
	// RemoveAll
	// Parameters:
	//  - locator interface{}
	//  a locator to remove reference
	// Returns interface{}
	// the removed component reference.
	Remove(locator interface{}) interface{}
	// 	Removes all component references that match the specified locator.
	// Parameters:
	//  - locator interface{}
	//  the locator to remove references by.
	// Returns []interface{}
	// a list, containing all removed references.
	RemoveAll(locator interface{}) []interface{}
	// 	Gets locators for all registered component references in this reference map.
	// Returns []interface{}
	// a list with component locators.
	GetAllLocators() []interface{}
	// Gets all component references registered in this reference map.
	// Returns []interface{}
	// a list with component references.
	GetAll() []interface{}
	// Gets all component references that match specified locator.
	// Parameters:
	//  - locator interface{}
	//  the locator to find references by.
	// Returns []interface{}
	// a list with matching component references or empty list if nothing was found.
	GetOptional(locator interface{}) []interface{}
	// 	Gets all component references that match specified locator. At least one component reference must be present. If it doesn't the method throws an error.
	// throws
	// a ReferenceException when no references found.
	// Parameters:
	//  - locator interface{}
	//  the locator to find references by.
	// Returns []interface{}
	// a list with matching component references.
	GetRequired(locator interface{}) ([]interface{}, error)
	// 	Gets an optional component reference that matches specified locator.
	// Parameters:
	//  - locator interface{}
	//  the locator to find references by.
	// Returns interface{}
	// a matching component reference or nil if nothing was found.
	GetOneOptional(locator interface{}) interface{}
	// Gets a required component reference that matches specified locator.
	// throws
	// a ReferenceError when no references found.
	// Parameters:
	//  - locator interface{}
	//  the locator to find a reference by.
	// Returns interface{}
	// a matching component reference.
	GetOneRequired(locator interface{}) (interface{}, error)
	// 	Gets all component references that match specified locator.
	// throws
	// a ReferenceError when required is set to true but no references found.
	// Parameters:
	// 	 - locator interface{}
	// 	 the locator to find a reference by.
	// 	 - required bool
	// 	 forces to raise an exception if no reference is found.
	// Returns []interface{}
	// a list with matching component references.
	Find(locator interface{}, required bool) ([]interface{}, error)
}
