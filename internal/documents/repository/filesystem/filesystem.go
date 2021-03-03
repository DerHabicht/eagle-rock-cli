package filesystem

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/pkg/errors"
)

type FileRepository struct {
}

func NewFileRepository() (FileRepository, error) {
	return FileRepository{}, errors.New("NewFileRepository() not implemented")
}

func (fr FileRepository) NewDocument(document.IDocument) error {
	return errors.New("NewDocument() not implemented")
}

func (fr FileRepository) LoadDocument(controlNumber string) document.IDocument {
	panic(errors.New("NewDocument() not implemented"))
}

func (fr FileRepository) LoadTemplate(controlNumber string) template.ITemplate {
	panic(errors.New("NewDocument() not implemented"))
}
