package convert

type TFloatConverter struct{}

var FloatConverter *TFloatConverter = &TFloatConverter{}

func (c *TFloatConverter) ToNullableFloat(value interface{}) *float32 {
	return ToNullableFloat(value)
}

func (c *TFloatConverter) ToFloat(value interface{}) float32 {
	return ToFloat(value)
}

func (c *TFloatConverter) ToFloatWithDefault(value interface{}, defaultValue float32) float32 {
	return ToFloatWithDefault(value, defaultValue)
}

func ToNullableFloat(value interface{}) *float32 {
	v := ToNullableDouble(value)
	if v == nil {
		return nil
	}
	r := float32(*v)
	return &r
}

func ToFloat(value interface{}) float32 {
	return ToFloatWithDefault(value, 0)
}

func ToFloatWithDefault(value interface{}, defaultValue float32) float32 {
	r := ToNullableFloat(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
