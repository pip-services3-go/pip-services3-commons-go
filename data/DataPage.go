package data

type DataPage struct {
	Total *int64        `json:"total"`
	Data  []interface{} `json:"data"`
}

func NewEmptyDataPage() *DataPage {
	return &DataPage{}
}

func NewDataPage(total *int64, data []interface{}) *DataPage {
	return &DataPage{Total: total, Data: data}
}
