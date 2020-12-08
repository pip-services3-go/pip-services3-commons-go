package run

/*
Interface for components that can be called to execute work.

Example:
 type EchoComponent {}
     ...
     func  (ec* EchoComponent) Execute(correlationId: string, args: Parameters) (result interface{}, err error) {
         return nil, result = args.getAsObject("message")
     }
 
 echo := EchoComponent{};
 message = "Test";
 res, err = echo.Execute("123", NewParametersFromTuples("message", message));
 fmt.Println(res);
*/
type IExecutable interface {
	// 	Executes component with arguments and receives execution result.
	// Parameters:
	//  - correlationId string
	//  transaction id to trace execution through call chain.
	//  - args *Parameters
	//  execution arguments.
	// Return interface{}, error
	// result or execution and error
	Execute(correlationId string, args *Parameters) (result interface{}, err error)
}
