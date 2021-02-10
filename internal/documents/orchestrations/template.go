package orchestrations

import "github.com/derhabicht/eagle-rock-cli/internal/documents/models"

type Template interface {
	// Inject performs the text replacements on the
	Inject(header models.Header, text []byte) ([]byte, error)
}
