package preprocessor

import (
	"github.com/pkg/errors"
	"os/exec"
)

type PandocPreprocessor struct {
	inputType string
	outputType string
}

func NewPandocPreprocessor(inputType string, outputType string) PandocPreprocessor {
	return PandocPreprocessor{
		inputType: inputType,
		outputType: outputType,
	}
}

func (pp PandocPreprocessor) Preprocess(text string) (string, error) {
	cmd := exec.Command("pandoc", "--from", pp.inputType, "--to", pp.outputType)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", errors.WithMessage(err, "failed to open a pipe to Pandoc")
	}

	go func() {
		defer stdin.Close()
		stdin.Write([]byte(text))
	}()

	output, err := cmd.Output()

	return string(output), err
}
