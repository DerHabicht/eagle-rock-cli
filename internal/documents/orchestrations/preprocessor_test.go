package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/preprocessors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockPreprocessor struct {
	mock.Mock
}

func (m *MockHeaderParser) Preprocess(text string) (string, error) {
	args := m.Called(text)
	return args.String(0), args.Error(1)
}

func TestPandocPreprocessor_IsPreprocessor(t *testing.T) {
	var _ Preprocessor = (*preprocessors.PandocPreprocessor)(nil)
	assert.True(t, true)
}

func TestNullPreprocessor_IsPreprocessor(t *testing.T) {
	var _ Preprocessor = (*preprocessors.NullPreprocessor)(nil)
	assert.True(t, true)
}
