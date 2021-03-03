package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/pkg/errors"
)

type Builder struct {
	SourceRepo      repository.IRepository
	DestinationRepo repository.IRepository
}

func (b Builder) Build() ([]byte, error) {
	// Step 1: Load document to be built
	// Step 2: Preprocess document body
	// Step 3: Load template
	// Step 4: Inject document content into template
	// Step 5: Build document
	// Step 6: Write build artifact

	return nil, errors.New("Build() method not implmented")
}
