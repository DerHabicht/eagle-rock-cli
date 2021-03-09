package template

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLatexTemplate_IsITemplate(t *testing.T) {
	var _ ITemplate = (*LatexTemplate)(nil)
	assert.True(t, true)
}

func TestLatexTemplate_Inject(t *testing.T) {
	input := map[string]interface{}{
		"FOO":  "bar",
		"GENE": []string{"ODBC!", "ODBC!", "ODBC!"},
	}

	testTemplate := `
Little bunny foo foo went to the %FOO%

Gene says the best things are:
%GENE%
`
	expected := `
Little bunny foo foo went to the bar

Gene says the best things are:
\gene{%
    \item ODBC!
    \item ODBC!
    \item ODBC!
}
`
	test := LatexTemplate{}
	test.Init(testTemplate)

	result, err := test.Inject(input)

	assert.NoError(t, err)
	assert.Equal(t, strings.Trim(expected, "\n"), strings.Trim(result, "\n"))
}
