package convert

import (
	"fmt"
	"strings"
	"time"
)

type TBooleanConverter struct{}

var BooleanConverter *TBooleanConverter = &TBooleanConverter{}

func (c *TBooleanConverter) ToNullableBoolean(value interface{}) *bool {
	return ToNullableBoolean(value)
}

func (c *TBooleanConverter) ToBoolean(value interface{}) bool {
	return ToBoolean(value)
}

func (c *TBooleanConverter) ToBooleanWithDefault(value interface{}, defaultValue bool) bool {
	return ToBooleanWithDefault(value, defaultValue)
}

func ToNullableBoolean(value interface{}) *bool {
	if value == nil {
		return nil
	}

	var v string

	switch value.(type) {
	case bool:
		r := value.(bool)
		return &r

	case string:
		v = strings.ToLower(value.(string))

	case time.Duration:
		d := value.(time.Duration)
		r := d.Nanoseconds() > 0
		return &r

	default:
		v = strings.ToLower(fmt.Sprint(value))
	}

	if v == "1" || v == "true" || v == "t" || v == "yes" || v == "y" {
		r := true
		return &r
	}

	if v == "0" || v == "false" || v == "f" || v == "no" || v == "n" {
		r := false
		return &r
	}

	return nil
}

func ToBoolean(value interface{}) bool {
	return ToBooleanWithDefault(value, false)
}

func ToBooleanWithDefault(value interface{}, defaultValue bool) bool {
	r := ToNullableBoolean(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
