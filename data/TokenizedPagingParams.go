package data

import "github.com/pip-services3-go/pip-services3-commons-go/convert"

/*
Data transfer object to pass paging parameters for queries.

The page is defined by two parameters:

the token parameter a starting point for a new page. The token shall be received from previous searches.
the take parameter sets how many items to return in a page.
additionally, the optional total parameter tells to return total number of items in the query.
Remember: not all implementations support the total parameter because its generation may lead to
severe performance implications.

Example:
 filter := NewFilterParamsFromTuples("type", "Type1");
 paging := NewTokenizedPagingParams("", 100);

 err, page = myDataClient.getDataByFilter(filter, paging);
*/
type TokenizedPagingParams struct {
	Token string `json:"token"`
	Take  *int64 `json:"take"`
	Total bool   `json:"total"`
}

// Creates a new instance.
// Returns *TokenizedPagingParams
func NewEmptyTokenizedPagingParams() *TokenizedPagingParams {
	return &TokenizedPagingParams{Token: "", Take: nil, Total: false}
}

// Creates a new instance and sets its values.
// Parameters:
//   - token string a token received from previous searches to define a starting point for this search.
//   - take interface{} the number of items to return.
//   - total interface{} true to return the total number of items.
// Returns *TokenizedPagingParams
func NewTokenizedPagingParams(token, take, total interface{}) *TokenizedPagingParams {
	c := TokenizedPagingParams{}

	c.Token = convert.StringConverter.ToString(token)
	c.Take = convert.LongConverter.ToNullableLong(take)
	c.Total = convert.BooleanConverter.ToBooleanWithDefault(total, false)

	return &c
}

// Gets the number of items to return in a page.
// Parameters:
//  - maxTake int64
//  the maximum number of items to return.
// Returns int64
// the number of items to return.
func (c *TokenizedPagingParams) GetTake(maxTake int64) int64 {
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

// Converts specified value into TokenizedPagingParams.
// Parameters:
//  - value interface{}
//  value to be converted
// Returns *PagingParams
// a newly created PagingParams.
func NewTokenizedPagingParamsFromValue(value interface{}) *TokenizedPagingParams {
	v, ok := value.(*TokenizedPagingParams)
	if ok {
		return v
	}

	maps := NewAnyValueMapFromValue(value)
	return NewTokenizedPagingParamsFromMap(maps)
}

// Creates a new TokenizedPagingParams from a list of key-value pairs called tuples.
// Parameters
//  - tuples ...interface{}
//  a list of values where odd elements are keys and the following even elements are values
// Returns *TokenizedPagingParams
// a newly created TokenizedPagingParams.
func NewTokenizedPagingParamsFromTuples(tuples ...interface{}) *TokenizedPagingParams {
	maps := NewAnyValueMapFromTuplesArray(tuples)
	return NewTokenizedPagingParamsFromMap(maps)
}

// Creates a new TokenizedPagingParams and sets it parameters from the specified map
// Parameters:
//  - value AnyValueMap
//  a AnyValueMap or StringValueMap to initialize this TokenizedPagingParams
// Returns *TokenizedPagingParams
// a newly created TokenizedPagingParams.
func NewTokenizedPagingParamsFromMap(value *AnyValueMap) *TokenizedPagingParams {
	c := TokenizedPagingParams{}

	c.Token = value.GetAsString("token")
	c.Take = value.GetAsNullableLong("take")
	c.Total = value.GetAsBooleanWithDefault("total", false)

	return &c
}
