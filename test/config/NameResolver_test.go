package test_config

import (
	"testing"

	conf "github.com/pip-services3-go/pip-services3-commons-go/v3/config"
	"github.com/stretchr/testify/assert"
)

func TestResolveName(t *testing.T) {
	var config = conf.NewConfigParamsFromTuples("id", "ABC")
	var name = conf.NameResolver.Resolve(config)
	assert.Equal(t, "ABC", name)

	config = conf.NewConfigParamsFromTuples("name", "ABC")
	name = conf.NameResolver.Resolve(config)
	assert.Equal(t, "ABC", name)
}

func TestResolveEmptyName(t *testing.T) {
	var config = conf.NewConfigParamsFromTuples()
	var name = conf.NameResolver.Resolve(config)
	assert.Equal(t, "", name)
}
