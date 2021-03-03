package repository

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
)

type IRepository interface {
	NewDocument(document.IDocument) error
	LoadDocument(controlNumber string) document.IDocument
	LoadTemplate(name string) template.ITemplate
}
