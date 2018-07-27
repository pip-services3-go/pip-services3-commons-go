package convert

import (
	"strconv"
	"time"
)

type TLongConverter struct{}

var LongConverter *TLongConverter = &TLongConverter{}

func (c *TLongConverter) ToNullableLong(value interface{}) *int64 {
	return ToNullableLong(value)
}

func (c *TLongConverter) ToLong(value interface{}) int64 {
	return ToLong(value)
}

func (c *TLongConverter) ToLongWithDefault(value interface{}, defaultValue int64) int64 {
	return ToLongWithDefault(value, defaultValue)
}

func ToNullableLong(value interface{}) *int64 {
	if value == nil {
		return nil
	}

	var r int64 = 0

	switch value.(type) {
	case int8:
		r = (int64)(value.(int8))
	case uint8:
		r = (int64)(value.(uint8))
	case int:
		r = (int64)(value.(int))
	case int16:
		r = (int64)(value.(int16))
	case uint16:
		r = (int64)(value.(uint16))
	case int32:
		r = (int64)(value.(int32))
	case uint32:
		r = (int64)(value.(uint32))
	case int64:
		r = (int64)(value.(int64))
	case uint64:
		r = (int64)(value.(uint64))
	case float32:
		r = (int64)(value.(float32))
	case float64:
		r = (int64)(value.(float64))

	case bool:
		v := value.(bool)
		if v == true {
			r = 1
		}

	case time.Time:
		r = value.(time.Time).Unix()

	case time.Duration:
		r = value.(time.Duration).Nanoseconds() / 1000000

	case string:
		v, ok := strconv.ParseFloat(value.(string), 0)
		if ok != nil {
			return nil
		}
		r = int64(v)

	default:
		return nil
	}

	return &r
}

func ToLong(value interface{}) int64 {
	return ToLongWithDefault(value, 0)
}

func ToLongWithDefault(value interface{}, defaultValue int64) int64 {
	r := ToNullableLong(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
