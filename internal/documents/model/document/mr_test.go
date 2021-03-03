package document

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMrHeader_IsIHeader(t *testing.T) {
	var _ IHeader = (*MrHeader)(nil)
	assert.True(t, true)
}

func TestMrSignature_IsISignature(t *testing.T) {
	var _ ISignature = (*MrSignature)(nil)
	assert.True(t, true)
}

func TestMr_IsIDocument(t *testing.T) {
	var _ IDocument = (*Mr)(nil)
	assert.True(t, true)
}
