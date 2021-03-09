package document

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoHeader_IsIHeader(t *testing.T) {
	var _ IHeader = (*MemoHeader)(nil)
	assert.True(t, true)
}

func TestMemoSignature_IsISignature(t *testing.T) {
	var _ ISignature = (*MemoSignature)(nil)
	assert.True(t, true)
}
