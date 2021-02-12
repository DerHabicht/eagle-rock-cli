package header_parsers

import "github.com/derhabicht/eagle-rock-cli/internal/documents/models"

type NullHeaderParser struct {}

func (nhp NullHeaderParser) ParseHeader(content []byte) (models.Header, []byte, error) {
	return nil, content, nil
}

func NewNullHeaderParser() NullHeaderParser {
	return NullHeaderParser{}
}