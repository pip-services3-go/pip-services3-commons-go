package data

/*
Data transfer object that is used to pass results of paginated queries. It contains items of retrieved
page and optional total number of items.
Most often this object type is used to send responses to paginated queries.
Pagination parameters are defined by TokenizedPagingParams object. The token parameter in the TokenizedPagingParams
there determines a starting point for a new search. It is received in the TokenizedDataPage from the previous search.
The takes parameter sets number of items to return in the page.
And the optional total parameter tells to return total number of items in the query.
Remember: not all implementations support the total parameter because its generation may lead to severe
performance implications.
see
TokenizedPagingParams

Example:
    err, page = myDataClient.getDataByFilter(
      "123",
      FilterParams.fromTuples("completed": true),
      NewTokenizedPagingParams("", 100, true)
    };

 	if err != nil {
 		panic()
 	}
 	for item range page.Data {
         fmt.Println(item);
     }
 );
*/
type TokenizedDataPage struct {
	Token string        `json:"token"`
	Data  []interface{} `json:"data"`
}

// Creates a new empty instance of data page.
// Returns *TokenizedDataPage
func NewEmptyTokenizedDataPage() *DataPage {
	return &DataPage{}
}

// Creates a new instance of data page and assigns its values.
// Parameters:
//  - token a token that defines a starting point for next search
//  - data []interface{}
//  a list of items from the retrieved page.
// Returns *TokenizedDataPage
func NewTokenizedDataPage(token string, data []interface{}) *TokenizedDataPage {
	return &TokenizedDataPage{Token: token, Data: data}
}
