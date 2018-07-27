package data

import "github.com/pip-services-go/pip-services-commons-go/convert"

type PagingParams struct {
	Skip  *int64
	Take  *int64
	Total bool
}

func NewEmptyPagingParams() *PagingParams {
	return &PagingParams{Skip: nil, Take: nil, Total: true}
}

func NewPagingParams(skip, take, total interface{}) *PagingParams {
	c := PagingParams{}

	c.Skip = convert.LongConverter.ToNullableLong(skip)
	c.Take = convert.LongConverter.ToNullableLong(take)
	c.Total = convert.BooleanConverter.ToBooleanWithDefault(total, true)

	return &c
}

func (c *PagingParams) GetSkip(minSkip int64) int64 {
	if c.Skip == nil {
		return minSkip
	}
	if *c.Skip < minSkip {
		return minSkip
	}
	return *c.Skip
}

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

func NewPagingParamsFromValue(value interface{}) *PagingParams {
	v, ok := value.(*PagingParams)
	if ok {
		return v
	}

	maps := NewAnyValueMapFromValue(value)
	return NewPagingParamsFromMap(maps)
}

func NewPagingParamsFromTuples(tuples ...interface{}) *PagingParams {
	maps := NewAnyValueMapFromTuplesArray(tuples)
	return NewPagingParamsFromMap(maps)
}

func NewPagingParamsFromMap(value *AnyValueMap) *PagingParams {
	c := PagingParams{}

	c.Skip = value.GetAsNullableLong("skip")
	c.Take = value.GetAsNullableLong("take")
	c.Total = value.GetAsBooleanWithDefault("total", true)

	return &c
}
