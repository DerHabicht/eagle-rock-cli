package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockBuilder struct {
	mock.Mock
}

func (m *MockBuilder) BuildDocument(text string) error {
	args := m.Called(text)
	return args.Error(0)
}

func TestPdflatexBuilder_IsBuilder(t *testing.T) {
	var _ Builder = (*services.PdflatexBuilder)(nil)
	assert.True(t, true)
}
