package data

type SortField struct {
	Name      string `json:"name"`
	Ascending bool   `json:"ascending"`
}

func NewEmptySortField() SortField {
	return SortField{}
}

func NewSortField(name string, ascending bool) SortField {
	return SortField{
		Name:      name,
		Ascending: ascending,
	}
}
