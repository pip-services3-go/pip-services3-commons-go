package test_reflect

type RootClass struct {
	rootPrivateField int
	RootPublicField  string
}

func (c *RootClass) RootMethod() {
}

type NestedClass struct {
	PublicField int
}

type TestClass struct {
	RootClass

	privateField int
	PublicField  string

	NestedField *NestedClass
}

func (c *TestClass) privateMethod() {
}

func (c *TestClass) PublicMethod(arg1 int, arg2 int) int {
	return arg1 + arg2
}
