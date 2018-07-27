package data

type FilterParams struct {
	StringValueMap
}

func NewEmptyFilterParams() *FilterParams {
	return &FilterParams{
		StringValueMap: *NewEmptyStringValueMap(),
	}
}

func NewFilterParams(values map[string]string) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromMaps(values),
	}
}

func NewFilterParamsFromValue(value interface{}) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromValue(value),
	}
}

func NewFilterParamsFromTuples(tuples ...interface{}) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromTuplesArray(tuples),
	}
}

func NewFilterParamsFromString(line string) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromString(line),
	}
}
