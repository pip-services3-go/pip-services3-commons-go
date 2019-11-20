package refer

/*
Interface for components that depends on other components.

If component requires explicit notification to unset references it shall additionally implement IUnreferenceable interface.

see
IReferences

see
IUnreferenceable

see
Referencer

Example
type MyController {
	_persistence IPersistence
}

    func (mc* MyController) setReferences(references IReferences) {
        mc._persistence = references.getOneRequired(
            NewDescriptor("mygroup", "persistence", "*", "*", "1.0"))
        );
    }
    ...
}
*/
type IReferenceable interface {
	// 	Sets references to dependent components.
	// see
	// IReferences
	// Parameters:
	// 			- references IReferences
	// 			references to locate the component dependencies.
	SetReferences(references IReferences)
}
