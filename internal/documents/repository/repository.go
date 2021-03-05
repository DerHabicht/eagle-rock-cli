package repository

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
)

type IRepository interface {
	NewDocument(document.IDocument) error
	LoadDocument(documents.ControlNumber) (document.IDocument, error)
	SaveCompiledDocument(documents.ControlNumber, []byte) error
	LoadTemplate(string, template.ITemplate) error
}
