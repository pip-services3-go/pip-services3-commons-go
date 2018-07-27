package data

type SortParams []SortField

func NewEmptySortParams() *SortParams {
	c := make(SortParams, 0, 10)
	return &c
}

func NewSortParams(fields []SortField) *SortParams {
	c := make(SortParams, len(fields))
	copy(c, fields)
	return &c
}
