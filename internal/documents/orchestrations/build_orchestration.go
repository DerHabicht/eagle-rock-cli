package orchestrations

import (
	"fmt"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/builders"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/content_readers"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/header_parsers"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/macro_processors"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/preprocessors"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/services/templates"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

type BuildOrchestration struct {
	reader         ContentReader
	headerParser   HeaderParser
	macroProcessor MacroProcessor
	preprocessor   Preprocessor
	template       Template
	builder        Builder
}

func NewBuildOrchestration (
	reader ContentReader,
	headerParser HeaderParser,
	macroProcessor MacroProcessor,
	preprocessor Preprocessor,
	template Template,
	builder Builder,
) BuildOrchestration {
	return BuildOrchestration{
		reader:         reader,
		headerParser:   headerParser,
		macroProcessor: macroProcessor,
		preprocessor:   preprocessor,
		template:       template,
		builder:        builder,
	}
}

func (o BuildOrchestration) Build(controlNumber string) error {
	year := parseYear(controlNumber)

	content, err := o.reader.Read(year, controlNumber)
	if err != nil {
		return errors.WithMessagef(err, "could not read %s", controlNumber)
	}

	header, text, err := o.headerParser.ParseHeader(content)
	if err != nil {
		return errors.WithMessagef(err, "failed to parse header for %s", controlNumber)
	}

	text, err = o.macroProcessor.ProcessMacros(text)
	if err != nil {
		return errors.WithMessagef(err, "failed to process macros in %s", controlNumber)
	}

	text, err = o.preprocessor.Preprocess(text)
	if err != nil {
		return errors.WithMessagef(err, "preprocessing of %s failed", controlNumber)
	}

	content, err = o.template.Inject(header, text)
	if err != nil {
		return errors.WithMessagef(err, "failed to inject content of %s into template", controlNumber)
	}

	err = o.builder.BuildDocument(year, controlNumber, content)
	if err != nil {
		return errors.WithMessagef(err, "failed to build %s", controlNumber)
	}

	return nil
}

// TODO: Do some data validation here
func parseYear(controlNumber string) string {
	i := strings.IndexRune(controlNumber, '-')
	return fmt.Sprintf("20%s", controlNumber[i+1:i+3])
}

func NewMemoForRecordBuildOrchestration() (BuildOrchestration, error) {
	contentDir, ok := viper.Get("content").(map[string]interface{})["mr_dir"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `content` is malformed: %#v", viper.Get("content"))
	}

	publishDir, ok := viper.Get("published").(map[string]interface{})["mr_dir"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `published` is malformed: %#v", viper.Get("published"))
	}

	templateDir, ok := viper.Get("latex_templates").(map[string]interface{})["mr"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `published` is malformed: %#v", viper.Get("pubilshed"))
	}

	texTemplate, err := ioutil.ReadFile(templateDir)
	if err != nil {
		return BuildOrchestration{}, errors.WithMessage(err, "Failed to read LaTeX template for Memoranda for Record.")
	}

	orch := NewBuildOrchestration(
		content_readers.NewFileContentReader(contentDir, ".md"),
		header_parsers.NewMemoForRecordHeaderParser(),
		macro_processors.NewNullMacroProcessor(),
		preprocessors.NewPandocPreprocessor("latex"),
		templates.NewLatexTemplate(texTemplate),
		builders.NewPdflatexBuilder(publishDir),
	)

	return orch, nil
}

func NewWarnoBuildOrchestration() (BuildOrchestration, error) {
	contentDir, ok := viper.Get("content").(map[string]interface{})["warno_dir"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `content` is malformed: %#v", viper.Get("content"))
	}

	publishDir, ok := viper.Get("published").(map[string]interface{})["warno_dir"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `published` is malformed: %#v", viper.Get("published"))
	}

	templateDir, ok := viper.Get("latex_templates").(map[string]interface{})["warno"].(string)
	if !ok {
		return BuildOrchestration{}, errors.Errorf("Configuration key `published` is malformed: %#v", viper.Get("pubilshed"))
	}

	texTemplate, err := ioutil.ReadFile(templateDir)
	if err != nil {
		return BuildOrchestration{}, errors.WithMessage(err, "Failed to read LaTeX template for Memoranda for Record.")
	}

	orch := NewBuildOrchestration(
		content_readers.NewFileContentReader(contentDir, ".txt"),
		header_parsers.NewNullHeaderParser(),
		macro_processors.NewNullMacroProcessor(),
		preprocessors.NewNullPreprocessor(),
		templates.NewLatexTemplate(texTemplate),
		builders.NewPdflatexBuilder(publishDir),
	)

	return orch, nil
}
