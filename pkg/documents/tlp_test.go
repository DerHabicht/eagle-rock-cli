package documents

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTlp_ValidStringWithCaveats(t *testing.T) {
	input := "TLP:RED//FOO/BAR"

	expected := Tlp{
		level:   RED,
		caveats: []string{"FOO", "BAR"},
	}

	result, err := ParseTlp(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseTlp_ValidStringWithoutCaveats(t *testing.T) {
	input := "TLP:AMBER"

	expected := Tlp{
		level:   AMBER,
		caveats: nil,
	}

	result, err := ParseTlp(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseTlp_ParsingIsNotCaseSensitive(t *testing.T) {
	input := "tlp:red//foo/bar"

	expected := Tlp{
		level:   RED,
		caveats: []string{"FOO", "BAR"},
	}

	result, err := ParseTlp(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseTlp_InvalidLevelCausesError(t *testing.T) {
	input := "CLASSIFIED"

	_, err := ParseTlp(input)

	assert.Error(t, err)
}

func TestParseTlp_MoreThanOneDoubleSlashCausesError(t *testing.T) {
	input := "TLP:RED//FOO//BAR"

	_, err := ParseTlp(input)

	assert.Error(t, err)
}

func TestParseTlp_OnlySingleSlashesCausesError(t *testing.T) {
	input := "TLP:RED/FOO/BAR"

	_, err := ParseTlp(input)

	assert.Error(t, err)
}
