package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPandocPreprocessor_Preprocess(t *testing.T) {
	test := `
####
This is a test memorandum.

####
For use in testing things.

####
By Testificates.
`

	expected := `\hypertarget{section}{%
\paragraph{}\label{section}}

This is a test memorandum.

\hypertarget{section-1}{%
\paragraph{}\label{section-1}}

For use in testing things.

\hypertarget{section-2}{%
\paragraph{}\label{section-2}}

By Testificates.
`

	p := NewPandocPreprocessor("latex")

	result, err := p.Preprocess([]byte(test))

	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}
