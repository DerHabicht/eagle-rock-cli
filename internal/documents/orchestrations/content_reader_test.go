package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockContentReader struct {
	mock.Mock
}

func (m *MockContentReader) Read(controlNumber string) ([]byte, error) {
	args := m.Called(controlNumber)
	return []byte(args.String(0)), args.Error(1)
}

func TestFileContentReader_IsContentReader(t *testing.T) {
	var _ ContentReader = (*services.FileContentReader)(nil)
	assert.True(t, true)
}
