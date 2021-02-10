package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoForRecordHeader_IsHeader(t *testing.T) {
	var _ Header = (*MemoForRecordHeader)(nil)
	assert.True(t, true)
}

func TestWarnoHeader_IsHeader(t *testing.T) {
	var _ Header = (*WarnoHeader)(nil)
	assert.True(t, true)
}

func TestOpordHeader_IsHeader(t *testing.T) {
	var _ Header = (*OpordHeader)(nil)
	assert.True(t, true)
}

func TestFragoHeader_IsHeader(t *testing.T) {
	var _ Header = (*FragoHeader)(nil)
	assert.True(t, true)
}
