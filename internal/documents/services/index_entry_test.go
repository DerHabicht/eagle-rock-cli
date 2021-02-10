package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFragoMeta_IsDocumentMeta(t *testing.T) {
	var _ IndexEntry = (*models.FragoIndexEntry)(nil)
	assert.True(t, true)
}

func TestMemoForRecordMeta_IsDocumentMeta(t *testing.T) {
	var _ IndexEntry = (*models.MemoForRecordIndexEntry)(nil)
	assert.True(t, true)
}

func TestOpordMeta_IsDocumentMeta(t *testing.T) {
	var _ IndexEntry = (*models.OpordIndexEntry)(nil)
	assert.True(t, true)
}

func TestWarno_IsDocumentMeta(t *testing.T) {
	var _ IndexEntry = (*models.WarnoIndexEntry)(nil)
	assert.True(t, true)
}
