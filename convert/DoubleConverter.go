package convert

import (
	"strconv"
	"time"
)

type TDoubleConverter struct{}

var DoubleConverter *TDoubleConverter = &TDoubleConverter{}

func (c *TDoubleConverter) ToNullableDouble(value interface{}) *float64 {
	return ToNullableDouble(value)
}

func (c *TDoubleConverter) ToDouble(value interface{}) float64 {
	return ToDouble(value)
}

func (c *TDoubleConverter) ToDoubleWithDefault(value interface{}, defaultValue float64) float64 {
	return ToDoubleWithDefault(value, defaultValue)
}

func ToNullableDouble(value interface{}) *float64 {
	if value == nil {
		return nil
	}

	var r float64 = 0

	switch value.(type) {
	case int8:
		r = (float64)(value.(int8))
	case uint8:
		r = (float64)(value.(uint8))
	case int:
		r = (float64)(value.(int))
	case int16:
		r = (float64)(value.(int16))
	case uint16:
		r = (float64)(value.(uint16))
	case int32:
		r = (float64)(value.(int32))
	case uint32:
		r = (float64)(value.(uint32))
	case int64:
		r = (float64)(value.(int64))
	case uint64:
		r = (float64)(value.(uint64))
	case float32:
		r = (float64)(value.(float32))
	case float64:
		r = (float64)(value.(float64))

	case bool:
		v := value.(bool)
		if v == true {
			r = 1.0
		}

	case time.Time:
		r = float64(value.(time.Time).Unix())

	case time.Duration:
		r = float64(value.(time.Duration).Nanoseconds() / 1000000)

	case string:
		var ok error
		r, ok = strconv.ParseFloat(value.(string), 0)
		if ok != nil {
			return nil
		}

	default:
		return nil
	}

	return &r
}

func ToDouble(value interface{}) float64 {
	return ToDoubleWithDefault(value, 0)
}

func ToDoubleWithDefault(value interface{}, defaultValue float64) float64 {
	r := ToNullableDouble(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
