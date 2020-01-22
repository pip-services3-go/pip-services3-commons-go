package data

import "github.com/pip-services3-go/pip-services3-commons-go/convert"

/*
Data transfer object to pass paging parameters for queries.

The page is defined by two parameters:

the skip parameter defines number of items to skip.
the take parameter sets how many items to return in a page.
additionally, the optional total parameter tells to return total number of items in the query.
Remember: not all implementations support the total parameter because its generation may lead to
severe performance implications.

Example:
filter := NewFilterParamsFromTuples("type", "Type1");
paging := NewPagingParams(0, 100);

err, page = myDataClient.getDataByFilter(filter, paging);
*/

type PagingParams struct {
	Skip  *int64
	Take  *int64
	Total bool
}

//Creates a new instance.
//Returns *PagingParams
func NewEmptyPagingParams() *PagingParams {
	return &PagingParams{Skip: nil, Take: nil, Total: false}
}

//Creates a new instance and sets its values.
//Parameters:
//			- skip interface{}
//			the number of items to skip.
//			- take interface{}
//			the number of items to return.
//			- total interface{}
//			true to return the total number of items.
//Returns *PagingParams
func NewPagingParams(skip, take, total interface{}) *PagingParams {
	c := PagingParams{}

	c.Skip = convert.LongConverter.ToNullableLong(skip)
	c.Take = convert.LongConverter.ToNullableLong(take)
	c.Total = convert.BooleanConverter.ToBooleanWithDefault(total, false)

	return &c
}

// Gets the number of items to skip.
// Parameters:
// 			- minSkip int64
// 			the minimum number of items to skip.
// Returns int64
// the number of items to skip.
func (c *PagingParams) GetSkip(minSkip int64) int64 {
	if c.Skip == nil {
		return minSkip
	}
	if *c.Skip < minSkip {
		return minSkip
	}
	return *c.Skip
}

// Gets the number of items to return in a page.
// Parameters:
// 			 - maxTake int64
// 			the maximum number of items to return.
// Returns int64
// the number of items to return.
func (c *PagingParams) GetTake(maxTake int64) int64 {
	if c.Take == nil {
		return maxTake
	}
	if *c.Take < 0 {
		return 0
	}
	if *c.Take > maxTake {
		return maxTake
	}
	return *c.Take
}

// Converts specified value into PagingParams.
// Parameters:
// 			- value interface{}
// 			value to be converted
// Returns *PagingParams
// a newly created PagingParams.
func NewPagingParamsFromValue(value interface{}) *PagingParams {
	v, ok := value.(*PagingParams)
	if ok {
		return v
	}

	maps := NewAnyValueMapFromValue(value)
	return NewPagingParamsFromMap(maps)
}

// Creates a new PagingParams from a list of key-value pairs called tuples.
// Parameters
// 			- tuples ...interface{}
// 			a list of values where odd elements are keys and the following even elements are values
// Returns *PagingParams
// a newly created PagingParams.
func NewPagingParamsFromTuples(tuples ...interface{}) *PagingParams {
	maps := NewAnyValueMapFromTuplesArray(tuples)
	return NewPagingParamsFromMap(maps)
}

// Creates a new PagingParams and sets it parameters from the specified map
// Parameters:
// 			value AnyValueMap
//			 a AnyValueMap or StringValueMap to initialize this PagingParams
// Returns *PagingParams
// a newly created PagingParams.
func NewPagingParamsFromMap(value *AnyValueMap) *PagingParams {
	c := PagingParams{}

	c.Skip = value.GetAsNullableLong("skip")
	c.Take = value.GetAsNullableLong("take")
	c.Total = value.GetAsBooleanWithDefault("total", false)

	return &c
}
