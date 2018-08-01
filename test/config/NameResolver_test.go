package convert

import (
	"testing"

	pconfig "github.com/pip-services-go/pip-services-commons-go/config"
	"github.com/stretchr/testify/assert"
)

func TestResolveName(t *testing.T) {
	var config = pconfig.NewConfigParamsFromTuples("id", "ABC")
	var name = pconfig.NameResolver.Resolve(config)
	assert.Equal(t, "ABC", name)

	config = pconfig.NewConfigParamsFromTuples("name", "ABC")
	name = pconfig.NameResolver.Resolve(config)
	assert.Equal(t, "ABC", name)
}

func TestResolveEmptyName(t *testing.T) {
	var config = pconfig.NewConfigParamsFromTuples()
	var name = pconfig.NameResolver.Resolve(config)
	assert.Equal(t, "", name)
}
