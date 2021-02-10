package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	uuid "github.com/satori/go.uuid"
)

type IndexEntry interface {
	DocumentUuid() uuid.UUID
	DocumentControlNumber() string
	DocumentStatusHistory() map[string]models.StatusHistoryEntry
}
