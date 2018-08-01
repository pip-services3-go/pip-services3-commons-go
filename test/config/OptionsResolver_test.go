package convert

import (
	"testing"

	pconfig "github.com/pip-services-go/pip-services-commons-go/config"
	"github.com/stretchr/testify/assert"
)

func TestResolveOptions(t *testing.T) {
	var config = pconfig.NewConfigParamsFromTuples(
		"test", "ABC",
		"options.test", "XYZ",
	)
	var options = pconfig.OptionsResolver.Resolve(config)
	assert.Equal(t, 1, options.Length())
	assert.Equal(t, "XYZ", options.GetAsString("test"))
}

func TestResolveOptionsWithDefault(t *testing.T) {
	var config = pconfig.NewConfigParamsFromTuples(
		"test", "ABC",
	)
	var options = pconfig.OptionsResolver.Resolve(config)
	assert.Equal(t, 0, options.Length())

	options = pconfig.OptionsResolver.ResolveWithDefault(config)
	assert.Equal(t, 1, options.Length())
	assert.Equal(t, "ABC", options.GetAsString("test"))
}
