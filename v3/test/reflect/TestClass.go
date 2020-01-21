package test_reflect

type RootClass struct {
	rootPrivateField int
	RootPublicField  string
	rootProperty     bool
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
	property     bool

	NestedField *NestedClass
}

func NewTestClass() *TestClass {
	return &TestClass{
		RootClass: RootClass{
			rootPrivateField: 111,
			RootPublicField:  "AAA",
			rootProperty:     true,
		},
		privateField: 222,
		PublicField:  "BBB",
		property:     true,
		NestedField: &NestedClass{
			PublicField: 333,
		},
	}
}

func (c *RootClass) rootPrivateProperty() {
}

func (c *RootClass) RootPublicProperty() bool {
	return c.rootProperty
}

func (c *RootClass) SetRootPublicProperty(value bool) {
	c.rootProperty = value
}

func (c *RootClass) RootPublicMethod(arg1 int, arg2 int) int {
	return arg1 + arg2
}

func (c *TestClass) PublicProperty() bool {
	return c.property
}

func (c *TestClass) SetPublicProperty(value bool) {
	c.property = value
}

func (c *TestClass) PublicMethod(arg1 int, arg2 int) int {
	return arg1 + arg2
}
