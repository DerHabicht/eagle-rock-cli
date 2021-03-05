package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/compiler"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/preprocessor"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
	"github.com/pkg/errors"
	"strings"
)

type Builder struct {
	repository   repository.IRepository
	preprocessor preprocessor.IPreprocessor
	template     template.ITemplate
	compiler     compiler.ICompiler
}

func NewBuilder(repo repository.IRepository, pre preprocessor.IPreprocessor, templ template.ITemplate, comp compiler.ICompiler) Builder {
	return Builder{
		repository:   repo,
		preprocessor: pre,
		template:     templ,
		compiler:     comp,
	}
}

func (b Builder) Build(controlNumber documents.ControlNumber) error {
	// Step 3: Load template
	templateName := strings.ToLower(controlNumber.Class.String())
	err := b.repository.LoadTemplate(templateName, b.template)
	if err != nil {
		return errors.WithMessagef(err, "failed to load build template for %s", controlNumber.String())
	}

	// Step 1: Load document to be built
	doc, err := b.repository.LoadDocument(controlNumber)
	if err != nil {
		return errors.WithMessagef(err, "failed to load document source for %s", controlNumber.String())
	}
	content := buildContentMap(doc)

	// Step 2: Preprocess document body
	body, err := b.preprocessor.Preprocess(doc.Body())
	content["BODY"] = body

	// Step 4: Inject document content into template
	src, err := b.template.Inject(content)
	if err != nil {
		return errors.WithMessagef(err, "failed to inject content into template for %s", controlNumber.String())
	}

	// Step 5: Build document
	artifact, err := b.compiler.Compile(src)
	if err != nil {
		return errors.WithMessagef(err, "failed to compile %s", controlNumber.String())
	}

	// Step 6: Write build artifact
	err = b.repository.SaveCompiledDocument(controlNumber, artifact)
	if err != nil {
		return errors.WithMessagef(err, "failed to save build artifact for %s", controlNumber.String())
	}

	return nil
}

func buildContentMap(document document.IDocument) map[string]interface{} {
	content := document.Header().HeaderFieldMap()
	sig := document.Signature().SignatureFieldMap()

	for k, v := range sig {
		content[k] = v
	}

	return content
}
