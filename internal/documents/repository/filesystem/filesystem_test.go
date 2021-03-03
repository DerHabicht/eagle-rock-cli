package filesystem

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepository_IsIRepository(t *testing.T) {
	var _ repository.IRepository = (*FileRepository)(nil)
	assert.True(t, true)
}
