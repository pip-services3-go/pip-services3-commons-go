package convert

import (
	"time"
)

type TDateTimeConverter struct{}

var DateTimeConverter *TDateTimeConverter = &TDateTimeConverter{}

func (c *TDateTimeConverter) ToNullableDateTime(value interface{}) *time.Time {
	return ToNullableDateTime(value)
}

func (c *TDateTimeConverter) ToDateTime(value interface{}) time.Time {
	return ToDateTime(value)
}

func (c *TDateTimeConverter) ToDateTimeWithDefault(value interface{}, defaultValue time.Time) time.Time {
	return ToDateTimeWithDefault(value, defaultValue)
}

func ToNullableDateTime(value interface{}) *time.Time {
	if value == nil {
		return nil
	}

	var r time.Time

	switch value.(type) {
	case int8:
		r = time.Unix((int64)(value.(int8)), 0)
	case uint8:
		r = time.Unix((int64)(value.(uint8)), 0)
	case int:
		r = time.Unix((int64)(value.(int)), 0)
	case int16:
		r = time.Unix((int64)(value.(int16)), 0)
	case uint16:
		r = time.Unix((int64)(value.(uint16)), 0)
	case int32:
		r = time.Unix((int64)(value.(int32)), 0)
	case uint32:
		r = time.Unix((int64)(value.(uint32)), 0)
	case int64:
		r = time.Unix((int64)(value.(int64)), 0)
	case uint64:
		r = time.Unix((int64)(value.(uint64)), 0)
	case float32:
		r = time.Unix((int64)(value.(float32)), 0)
	case float64:
		r = time.Unix((int64)(value.(float64)), 0)

	case time.Time:
		r = value.(time.Time)

	case string:
		v := value.(string)
		var ok error
		r, ok = time.Parse(time.RFC3339, v)
		if ok != nil {
			r, ok = time.Parse(time.RFC3339Nano, v)
		}
		if ok != nil {
			return nil
		}

	default:
		return nil
	}

	return &r
}

func ToDateTime(value interface{}) time.Time {
	return ToDateTimeWithDefault(value, time.Time{})
}

func ToDateTimeWithDefault(value interface{}, defaultValue time.Time) time.Time {
	r := ToNullableDateTime(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
