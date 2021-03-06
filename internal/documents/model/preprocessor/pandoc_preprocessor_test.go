package preprocessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatexTemplate_IsIPreprocessor(t *testing.T) {
	var _ IPreprocessor = (*PandocPreprocessor)(nil)
	assert.True(t, true)
}

