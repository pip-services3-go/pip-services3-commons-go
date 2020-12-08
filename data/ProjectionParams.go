package data

import (
	"strings"
)

/*
Defines projection parameters with list if fields to include into query results.
The parameters support two formats: dot format and nested format.
The dot format is the standard way to define included fields and subfields
using dot object notation: "field1,field2.field21,field2.field22.field221".
As alternative the nested format offers a more compact representation: "field1,field2(field21,field22(field221))".

Example:
 filter := NewFilterParamsFromTuples("type", "Type1");
 paging := NewPagingParams(0, 100);
 projection = NewProjectionParamsFromString("field1,field2(field21,field22)")

 err, page := myDataClient.getDataByFilter(filter, paging, projection);
*/
type ProjectionParams struct {
	values []string
}

// Creates a new instance of the projection parameters and assigns its value.
// Returns *ProjectionParams
func NewEmptyProjectionParams() *ProjectionParams {
	return &ProjectionParams{
		values: make([]string, 0, 10),
	}
}

// Creates a new instance of the projection parameters and assigns its from string value.
// Parameters:
//  - values []string
// Returns *ProjectionParams
func NewProjectionParamsFromStrings(values []string) *ProjectionParams {
	c := &ProjectionParams{
		values: make([]string, len(values)),
	}
	copy(c.values, values)
	return c
}

// Creates a new instance of the projection parameters and assigns its from AnyValueArray values.
// Parameters:
//  - values *AnyValueArray
// Returns *ProjectionParams
func NewProjectionParamsFromAnyArray(values *AnyValueArray) *ProjectionParams {
	if values == nil {
		return NewEmptyProjectionParams()
	}

	c := &ProjectionParams{
		values: make([]string, 0, values.Len()),
	}

	for index := 0; index < values.Len(); index++ {
		value := values.GetAsString(index)
		if value != "" {
			c.values = append(c.values, value)
		}
	}

	return c
}

// Return raw values []string
func (c *ProjectionParams) Value() []string {
	return c.values
}

// Gets or sets the length of the array. This is a number one
// higher than the highest element defined in an array.
func (c *ProjectionParams) Len() int {
	return len(c.values)
}

// Get value by index
// Parameters:
//  - index int
//  an index of element
// Return string
func (c *ProjectionParams) Get(index int) string {
	return c.values[index]
}

// Set value in index position
// Parameters:
//  - index int
//  an index of element
//  - value string
//  value
func (c *ProjectionParams) Put(index int, value string) {
	if cap(c.values)+1 < index {
		a := make([]string, index+1, (index+1)*2)
		copy(a, c.values)
		c.values = a
	}

	c.values[index] = value
}

// Remove element by index
// Parameters:
//  - index int
//  an index of remove element
func (c *ProjectionParams) Remove(index int) {
	c.values = append(c.values[:index], c.values[index+1:]...)
}

// Appends new element to an array.
// Parameters:
//  - value string
func (c *ProjectionParams) Push(value string) {
	c.values = append(c.values, value)
}

// Appends new elements to an array.
// Parameters:
//  - value []string
func (c *ProjectionParams) Append(elements []string) {
	if elements != nil {
		c.values = append(c.values, elements...)
	}
}

// Clear elements
func (c *ProjectionParams) Clear() {
	c.values = make([]string, 0, 10)
}

// Returns a string representation of an array.
// Returns string
func (c *ProjectionParams) String() string {
	builder := ""

	for index := 0; index < c.Len(); index++ {
		if index > 0 {
			builder = builder + ","
		}

		builder = builder + c.Get(index)
	}

	return builder
}

// Converts specified value into ProjectionParams.
// see
// AnyValueArray.fromValue
// Parameters:
//  - value interface{}
//  value to be converted
// Returns *ProjectionParams
// a newly created ProjectionParams.
func NewProjectionParamsFromValue(value interface{}) *ProjectionParams {
	values := NewAnyValueArrayFromValue(value)
	return NewProjectionParamsFromAnyArray(values)
}

// Create new ProjectionParams and set values from values
// Parameters:
//  - values ...string
//  an values to parce
// Return *ProjectionParams
func ParseProjectionParams(values ...string) *ProjectionParams {
	c := NewEmptyProjectionParams()

	for index := 0; index < len(values); index++ {
		parseProjectionParamValue("", c, values[index])
	}

	return c
}

// Add parce value into exist ProjectionParams and add prefix
// Parameters:
//  - prefix string
//  prefix value
//  - c *ProjectionParams
//  ProjectionParams instance wheare need to add value
//  - value string
//  an values to parce
func parseProjectionParamValue(prefix string, c *ProjectionParams, value string) {
	if value != "" {
		value = strings.Trim(value, " \t\n\r")
	}

	openBracket := 0
	openBracketIndex := -1
	closeBracketIndex := -1
	commaIndex := -1

	breakCycleRequired := false
	for index := 0; index < len(value); index++ {
		switch value[index] {
		case '(':
			if openBracket == 0 {
				openBracketIndex = index
			}

			openBracket++
			break
		case ')':
			openBracket--

			if openBracket == 0 {
				closeBracketIndex = index

				if openBracketIndex >= 0 && closeBracketIndex > 0 {
					previousPrefix := prefix

					if prefix != "" {
						prefix = prefix + "." + value[:openBracketIndex]
					} else {
						prefix = value[:openBracketIndex]
					}

					subValue := value[openBracketIndex+1 : closeBracketIndex]
					parseProjectionParamValue(prefix, c, subValue)

					subValue = value[closeBracketIndex+1:]
					parseProjectionParamValue(previousPrefix, c, subValue)
					breakCycleRequired = true
				}
			}
			break
		case ',':
			if openBracket == 0 {
				commaIndex = index

				subValue := value[0:commaIndex]

				if subValue != "" {
					if prefix != "" {
						c.Push(prefix + "." + subValue)
					} else {
						c.Push(subValue)
					}
				}

				subValue = value[commaIndex+1:]

				if subValue != "" {
					parseProjectionParamValue(prefix, c, subValue)
					breakCycleRequired = true
				}
			}
			break
		}

		if breakCycleRequired {
			break
		}
	}

	if value != "" && openBracketIndex == -1 && commaIndex == -1 {
		if prefix != "" {
			c.Push(prefix + "." + value)
		} else {
			c.Push(value)
		}
	}
}
