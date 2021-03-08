package repository

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/artifact"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-lib/lib"
)

type IRepository interface {
	NewDocument(document.IDocument) error
	LoadDocument(lib.ControlNumber) (document.IDocument, error)
	SaveCompiledDocument(lib.ControlNumber, artifact.BuildArtifact) (string, error)
	LoadTemplate(string, template.ITemplate) error
}
