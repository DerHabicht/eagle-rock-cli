package builders

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/rwestlund/gotex"
	"io/ioutil"
	"path/filepath"
)

type PdflatexBuilder struct {
	publishDir string
}

func NewPdflatexBuilder(publishDir string) PdflatexBuilder {
	return PdflatexBuilder{
		publishDir: publishDir,
	}
}

func (pb PdflatexBuilder) BuildDocument(year string, controlNumber string, content []byte) error {
	pdf, err := gotex.Render(string(content), gotex.Options{})
	if err != nil {
		log.Debug().Str("document", string(content)).Msg("Document build failed, dumping content.")
		return errors.WithMessagef(err, "failed to build %s", controlNumber)
	}

	err = ioutil.WriteFile(filepath.Join(pb.publishDir, year, fmt.Sprintf("%s.pdf", controlNumber)), pdf, 0644)
	if err != nil {
		return errors.WithMessagef(err, "failed to write %s", controlNumber)
	}

	return nil
}
