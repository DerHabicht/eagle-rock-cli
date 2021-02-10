package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockHeaderParser struct {
	mock.Mock
}
func (m *MockHeaderParser) ParseHeader(raw []byte) (models.Header, string, error) {
	args := m.Called(raw)
	return args.Get(0).(models.Header), args.String(1), args.Error(2)
}

func TestMemoForRecordHeaderParser_IsHeaderParser(t *testing.T) {
	var _ HeaderParser = (*services.MemoForRecordHeaderParser)(nil)
	assert.True(t, true)
}
