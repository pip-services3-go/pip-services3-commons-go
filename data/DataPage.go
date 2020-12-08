package data

/*
Data transfer object that is used to pass results of paginated queries. It contains items of retrieved
page and optional total number of items.
Most often this object type is used to send responses to paginated queries.
Pagination parameters are defined by PagingParams object. The skip parameter in the PagingParams
there means how many items to skip. The takes parameter sets number of items to return in the page.
And the optional total parameter tells to return total number of items in the query.
Remember: not all implementations support the total parameter because its generation may lead to severe
performance implications.
see
PagingParams

Example:
 err, page = myDataClient.getDataByFilter(
     "123",
     FilterParams.fromTuples("completed": true),
     NewPagingParams(0, 100, true)
 	};
 
 	if err != nil {
 		panic()
 	}
 	for item range page.Data {
         fmt.Println(item);
     }
 );
*/
type DataPage struct {
	Total *int64        `json:"total"`
	Data  []interface{} `json:"data"`
}

// Creates a new empty instance of data page.
// Returns *DataPage
func NewEmptyDataPage() *DataPage {
	return &DataPage{}
}

// Creates a new instance of data page and assigns its values.
// Parameters:
//  - value data []interface{}
//  a list of items from the retrieved page.
//  - total int64
// Returns *DataPage
func NewDataPage(total *int64, data []interface{}) *DataPage {
	return &DataPage{Total: total, Data: data}
}
