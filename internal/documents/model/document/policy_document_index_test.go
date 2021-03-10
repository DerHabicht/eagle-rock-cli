package document

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicyDocumentIndex_IsIDocumentIndex(t *testing.T) {
	var _ IDocumentIndex = (*PolicyDocumentIndex)(nil)
	assert.True(t, true)
}
