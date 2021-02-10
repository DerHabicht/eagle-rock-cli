package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrago_IsDocument(t *testing.T) {
	var _ Document = (*models.Frago)(nil)
	assert.True(t, true)
}

func TestMemoForRecord_IsDocument(t *testing.T) {
	var _ Document = (*models.MemoForRecord)(nil)
	assert.True(t, true)
}

func TestOpord_IsDocument(t *testing.T) {
	var _ Document = (*models.Opord)(nil)
	assert.True(t, true)
}

func TestWarno_IsDocument(t *testing.T) {
	var _ Document = (*models.Warno)(nil)
	assert.True(t, true)
}
