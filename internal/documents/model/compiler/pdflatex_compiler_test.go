package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPdflatexCompiler_IsICompiler(t *testing.T) {
	var _ ICompiler = (*PdflatexCompiler)(nil)
	assert.True(t, true)
}
