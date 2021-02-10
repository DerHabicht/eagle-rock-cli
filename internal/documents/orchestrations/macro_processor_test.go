package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockMacroProcessor struct {
	mock.Mock
}

func (m *MockMacroProcessor) ProcessMacros(text string) (string, error) {
	args := m.Called(text)
	return args.String(0), args.Error(1)
}

func TestNullMacroProcessor_IsMacroProcessor(t *testing.T) {
	var _ MacroProcessor = (*services.NullMacroProcessor)(nil)
	assert.True(t, true)
}
