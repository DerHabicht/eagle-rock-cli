package services

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/compiler"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/preprocessor"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/template"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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

func (b Builder) Build(controlNumber lib.ControlNumber) error {
	log.Info().Msgf("Loading %s build template...", controlNumber.Class.String())
	templateName := strings.ToLower(controlNumber.Class.String())
	err := b.repository.LoadTemplate(templateName, b.template)
	if err != nil {
		return errors.WithMessagef(err, "failed to load build template for %s", controlNumber.String())
	}

	log.Info().Msgf("Loading document %s...", controlNumber.String())
	doc, err := b.repository.LoadDocument(controlNumber)
	if err != nil {
		return errors.WithMessagef(err, "failed to load document source for %s", controlNumber.String())
	}
	content := buildContentMap(doc)

	log.Info().Msg("Preprocessing document body...")
	body, err := b.preprocessor.Preprocess(doc.Body())
	content["BODY"] = body

	log.Info().Msg("Injecting content into build template...")
	src, err := b.template.Inject(content)
	if err != nil {
		return errors.WithMessagef(err, "failed to inject content into template for %s", controlNumber.String())
	}

	log.Info().Msg("Building document...")
	artifact, err := b.compiler.Compile(src)
	if err != nil {
		return errors.WithMessagef(err, "failed to compile %s", controlNumber.String())
	}

	log.Info().Msg("Saving document...")
	path, err := b.repository.SaveCompiledDocument(controlNumber, artifact)
	if err != nil {
		return errors.WithMessagef(err, "failed to save build artifact for %s", controlNumber.String())
	}

	log.Info().Msgf("Success! Document has been saved under: %s", path)
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
