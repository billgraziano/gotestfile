package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	assert := assert.New(t)
	m, err := Module()
	assert.NoError(err)
	assert.Equal("github.com/billgraziano/gotestfile", m)
}
