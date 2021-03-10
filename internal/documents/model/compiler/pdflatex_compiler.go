package compiler

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/artifact"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/rwestlund/gotex"
)

type PdflatexCompiler struct {
	options gotex.Options
}

func NewPdfLatexCompiler(texinputs string) PdflatexCompiler {
	return PdflatexCompiler{
		options: gotex.Options{
			Texinputs: texinputs,
		},
	}
}

func (pb PdflatexCompiler) Compile(source string) (artifact.BuildArtifact, error) {
	pdf, err := gotex.Render(source, pb.options)
	if err != nil {
		log.Debug().Str("document", source).Msg("Document build failed, dumping content.")
		return artifact.BuildArtifact{}, errors.WithStack(err)
	}

	return artifact.NewBuildArtifact("pdf", pdf), nil
}
