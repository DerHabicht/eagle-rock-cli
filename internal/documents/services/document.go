package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	uuid "github.com/satori/go.uuid"
)

type Document interface {
	Uuid() uuid.UUID
	ControlNumber() string
	DocumentType() models.DocumentType
	Header() models.Header
}
