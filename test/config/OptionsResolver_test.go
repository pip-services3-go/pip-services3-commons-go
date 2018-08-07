package test_config

import (
	"testing"

	conf "github.com/pip-services-go/pip-services-commons-go/config"
	"github.com/stretchr/testify/assert"
)

func TestResolveOptions(t *testing.T) {
	var config = conf.NewConfigParamsFromTuples(
		"test", "ABC",
		"options.test", "XYZ",
	)
	var options = conf.OptionsResolver.Resolve(config)
	assert.Equal(t, 1, options.Length())
	assert.Equal(t, "XYZ", options.GetAsString("test"))
}

func TestResolveOptionsWithDefault(t *testing.T) {
	var config = conf.NewConfigParamsFromTuples(
		"test", "ABC",
	)
	var options = conf.OptionsResolver.Resolve(config)
	assert.Equal(t, 0, options.Length())

	options = conf.OptionsResolver.ResolveWithDefault(config)
	assert.Equal(t, 1, options.Length())
	assert.Equal(t, "ABC", options.GetAsString("test"))
}
