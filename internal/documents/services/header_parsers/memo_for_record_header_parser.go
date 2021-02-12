package header_parsers

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"regexp"
)

type MemoForRecordHeaderParser struct {}

func NewMemoForRecordHeaderParser() MemoForRecordHeaderParser {
	return MemoForRecordHeaderParser{}
}

func (mfrhp MemoForRecordHeaderParser) ParseHeader(content []byte) (models.Header, []byte, error) {
	var header models.MemoForRecordHeader
	re := regexp.MustCompile(`(?ms)-{3}\n(.+)\n\.{3}(.*)`)

	b := re.FindSubmatch(content)

	err := yaml.Unmarshal(b[1], &header)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return header, b[2], nil
}
