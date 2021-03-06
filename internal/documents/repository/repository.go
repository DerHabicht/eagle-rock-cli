package repository

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/artifact"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
)

type IRepository interface {
	NewDocument(document.IDocument) error
	LoadDocument(documents.ControlNumber) (document.IDocument, error)
	SaveCompiledDocument(documents.ControlNumber, artifact.BuildArtifact) (string, error)
	LoadTemplate(string, template.ITemplate) error
}
