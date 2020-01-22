package test_convert

import (
	"testing"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/convert"
	"github.com/stretchr/testify/assert"
)

func TestToDateTime(t *testing.T) {
	assert.Nil(t, convert.ToNullableDateTime(nil))

	date1 := time.Date(1975, time.April, 8, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, date1, convert.ToDateTimeWithDefault(nil, date1))
	assert.Equal(t, date1, convert.ToDateTime(date1))
	assert.Equal(t, date1, convert.ToDateTime("1975-04-08T00:00:00Z"))
	assert.Equal(t, date1, convert.ToDateTime("1975-04-08T00:00:00.00Z"))

	date2 := time.Unix(123, 0)
	assert.Equal(t, date2, convert.ToDateTime(123))
	assert.Equal(t, date2, convert.ToDateTime(123.456))
}
