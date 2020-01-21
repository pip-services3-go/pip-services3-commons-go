package data

/*
Data transfer object used to pass filter parameters as simple key-value pairs.
see
StringValueMap

Example:
 filter := NewFilterParamsFromTuples(
    "type", "Type1",
    "from_create_time", new Date(2000, 0, 1),
    "to_create_time", new Date(),
    "completed", true
);
paging = NewPagingParams(0, 100);

err, page = myDataClient.getDataByFilter(filter, paging);
*/

type FilterParams struct {
	StringValueMap
}

// Creates a new instance.
// Returns *FilterParams
func NewEmptyFilterParams() *FilterParams {
	return &FilterParams{
		StringValueMap: *NewEmptyStringValueMap(),
	}
}

// Creates a new instance and initalizes it with elements from the specified map.
// Parameters:
// 			- value map[string]string
// 			a map to initialize this instance.
// Returns *FilterParams
func NewFilterParams(values map[string]string) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromMaps(values),
	}
}

// Converts specified value into FilterParams.
// Parameters:
// 			- value interface
// value to be converted
// Returns FilterParams
// a newly created FilterParams.
func NewFilterParamsFromValue(value interface{}) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromValue(value),
	}
}

// Creates a new FilterParams from a list of key-value pairs called tuples.
// Parameters:
// 			- tuples ...interface{}
// 			a list of values where odd elements are keys and the following even elements are values
// Returns *FilterParams
// a newly created FilterParams.
func NewFilterParamsFromTuples(tuples ...interface{}) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromTuplesArray(tuples),
	}
}

// Parses semicolon-separated key-value pairs and returns them as a FilterParams.
// see
// StringValueMap.FromString
// Parameters:
// 			- line string
// 			semicolon-separated key-value list to initialize FilterParams.
// Returns *FilterParams
func NewFilterParamsFromString(line string) *FilterParams {
	return &FilterParams{
		StringValueMap: *NewStringValueMapFromString(line),
	}
}
