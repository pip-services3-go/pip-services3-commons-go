package convert

import (
	"encoding/json"
)

func FromJson(value string) (interface{}, error) {
	if value == "" {
		return nil, nil
	}

	var m interface{}
	if err := json.Unmarshal([]byte(value), &m); err != nil {
		return nil, err
	}
	return m, nil
}

func ToJson(value interface{}) (string, error) {
	if value == nil {
		return "", nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

type TJsonConverter struct{}

var JsonConverter *TJsonConverter = &TJsonConverter{}

func (c *TJsonConverter) ToNullableMap(value string) *map[string]interface{} {
	return JsonToNullableMap(value)
}

func (c *TJsonConverter) ToMap(value string) map[string]interface{} {
	return JsonToMap(value)
}

func (c *TJsonConverter) ToMapWithDefault(value string, defaultValue map[string]interface{}) map[string]interface{} {
	return JsonToMapWithDefault(value, defaultValue)
}

func JsonToNullableMap(value string) *map[string]interface{} {
	v, _ := FromJson(value)
	if v == nil {
		return nil
	}
	return ToNullableMap(v)
}

func JsonToMap(value string) map[string]interface{} {
	return JsonToMapWithDefault(value, map[string]interface{}{})
}

func JsonToMapWithDefault(value string, defaultValue map[string]interface{}) map[string]interface{} {
	if m := JsonToNullableMap(value); m != nil {
		return *m
	}
	return map[string]interface{}{}
}
