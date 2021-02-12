package macro_processors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullMacroProcessor_ProcessMacros(t *testing.T) {
	test := "%Ignores% $all ${possible} #macro !<marks>"
	expected := test

	p := NewNullMacroProcessor()

	result, err := p.ProcessMacros([]byte(test))

	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}
