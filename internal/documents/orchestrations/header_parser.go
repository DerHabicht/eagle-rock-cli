package orchestrations

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
)

type HeaderParser interface {
	// ParseHeader processes the YAML document header inside the content.
	// It returns the associated header object, and the content body with the header document stripped.
	ParseHeader(content []byte) (models.Header, []byte, error)
}
