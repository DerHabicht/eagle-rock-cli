package services

import (
	"github.com/pkg/errors"
	"os/exec"
)

type PandocPreprocessor struct {
	outputType string
}

func NewPandocPreprocessor(outputType string) PandocPreprocessor {
	return PandocPreprocessor{
		outputType: outputType,
	}
}

func (pp PandocPreprocessor) Preprocess(text []byte) ([]byte, error) {
	cmd := exec.Command("pandoc", "--to", pp.outputType)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to open a pipe to Pandoc")
	}

	go func() {
		defer stdin.Close()
		stdin.Write(text)
	}()

	return cmd.Output()
}
