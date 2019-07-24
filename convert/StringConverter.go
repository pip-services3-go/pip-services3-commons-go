package convert

import (
	"fmt"
	refl "reflect"
	"strconv"
	"time"
)

type TStringConverter struct{}

var StringConverter *TStringConverter = &TStringConverter{}

func (c *TStringConverter) ToNullableString(value interface{}) *string {
	return ToNullableString(value)
}

func (c *TStringConverter) ToString(value interface{}) string {
	return ToString(value)
}

func (c *TStringConverter) ToStringWithDefault(value interface{}, defaultValue string) string {
	return ToStringWithDefault(value, defaultValue)
}

func ToNullableString(value interface{}) *string {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case string:
		r := value.(string)
		return &r

	case byte, uint, uint32, uint64, int, int32, int64:
		r := strconv.FormatInt(ToLong(value), 10)
		return &r

	case float32, float64:
		r := strconv.FormatFloat(ToDouble(value), 'f', -1, 64)
		return &r

	case bool:
		r := "false"
		if value.(bool) {
			r = "true"
		}
		return &r

	case time.Time:
		r := value.(time.Time).Format(time.RFC3339)
		return &r

	case time.Duration:
		r := strconv.FormatInt(value.(time.Duration).Nanoseconds()/1000000, 10)
		return &r

	default:
		val := refl.ValueOf(value)
		if val.Kind() == refl.Slice || val.Kind() == refl.Array {
			r := ""
			for index := 0; index < val.Len(); index++ {
				if len(r) > 0 {
					r += ","
				}
				r += fmt.Sprint(val.Index(index).Interface())
			}
			return &r
		}

		r := fmt.Sprint(value)
		return &r
	}
}

func ToString(value interface{}) string {
	return ToStringWithDefault(value, "")
}

func ToStringWithDefault(value interface{}, defaultValue string) string {
	r := ToNullableString(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
