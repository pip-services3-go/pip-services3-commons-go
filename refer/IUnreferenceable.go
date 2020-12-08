package refer

/*
Interface for components that require explicit clearing of references to dependent components.

see
IReferences

see
IReferenceable

Example
 type MyController  {
     _persistence IMyPersistence;
 }
 func (mc* MyController) setReferences(references *IReferences) {
     mc._persistence = references.GetOneRequired(
         NewDescriptor("mygroup", "persistence", "*", "*", "1.0")
     );
 }

 func (mc* MyController) UnsetReferences() {
     mc._persistence = nil;
 }


*/

type IUnreferenceable interface {
	//Unsets (clears) previously set references to dependent components.
	UnsetReferences()
}
