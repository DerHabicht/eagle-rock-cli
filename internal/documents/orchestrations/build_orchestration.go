package orchestrations

import (
	"github.com/pkg/errors"
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
	content, err := o.reader.Read(controlNumber)
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

	err = o.builder.BuildDocument(controlNumber, content)
	if err != nil {
		return errors.WithMessagef(err, "failed to build %s", controlNumber)
	}

	return nil
}
