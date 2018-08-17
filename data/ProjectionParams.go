package data

import (
	"strings"
)

type ProjectionParams struct {
	values []string
}

func NewEmptyProjectionParams() *ProjectionParams {
	return &ProjectionParams{
		values: make([]string, 0, 10),
	}
}

func NewProjectionParamsFromStrings(values []string) *ProjectionParams {
	c := &ProjectionParams{
		values: make([]string, len(values)),
	}
	copy(c.values, values)
	return c
}

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

func (c *ProjectionParams) Value() []string {
	return c.values
}

func (c *ProjectionParams) Len() int {
	return len(c.values)
}

func (c *ProjectionParams) Get(index int) string {
	return c.values[index]
}

func (c *ProjectionParams) Put(index int, value string) {
	if cap(c.values)+1 < index {
		a := make([]string, index+1, (index+1)*2)
		copy(a, c.values)
		c.values = a
	}

	c.values[index] = value
}

func (c *ProjectionParams) Remove(index int) {
	c.values = append(c.values[:index], c.values[index+1:]...)
}

func (c *ProjectionParams) Push(value string) {
	c.values = append(c.values, value)
}

func (c *ProjectionParams) Append(elements []string) {
	if elements != nil {
		c.values = append(c.values, elements...)
	}
}

func (c *ProjectionParams) Clear() {
	c.values = make([]string, 0, 10)
}

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

func NewProjectionParamsFromValue(value interface{}) *ProjectionParams {
	values := NewAnyValueArrayFromValue(value)
	return NewProjectionParamsFromAnyArray(values)
}

func ParseProjectionParams(values ...string) *ProjectionParams {
	c := NewEmptyProjectionParams()

	for index := 0; index < len(values); index++ {
		parseProjectionParamValue("", c, values[index])
	}

	return c
}

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
