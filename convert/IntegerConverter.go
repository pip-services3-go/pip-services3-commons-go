package convert

type TIntegerConverter struct{}

var IntegerConverter *TIntegerConverter = &TIntegerConverter{}

func (c *TIntegerConverter) ToNullableInteger(value interface{}) *int {
	return ToNullableInteger(value)
}

func (c *TIntegerConverter) ToInteger(value interface{}) int {
	return ToInteger(value)
}

func (c *TIntegerConverter) ToIntegerWithDefault(value interface{}, defaultValue int) int {
	return ToIntegerWithDefault(value, defaultValue)
}

func ToNullableInteger(value interface{}) *int {
	v := ToNullableLong(value)
	if v == nil {
		return nil
	}
	r := int(*v)
	return &r
}

func ToInteger(value interface{}) int {
	return ToIntegerWithDefault(value, 0)
}

func ToIntegerWithDefault(value interface{}, defaultValue int) int {
	r := ToNullableInteger(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
