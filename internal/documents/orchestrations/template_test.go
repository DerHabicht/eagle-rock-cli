package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockTemplate struct {
	mock.Mock
}

func (m *MockTemplate) Inject(header models.Header, content string) (string, error) {
	args := m.Called(header, content)
	return args.String(0), args.Error(1)
}

func TestLatexTemplate_IsTemplate(t *testing.T) {
	var _ Template = (*templates.LatexTemplate)(nil)
	assert.True(t, true)
}
