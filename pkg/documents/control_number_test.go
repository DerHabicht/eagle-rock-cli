package documents

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseControlNumber_WithoutSubsequence(t *testing.T) {
	input := "MR-20-001"

	expected := ControlNumber{
		Class:        MR,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  nil,
	}

	result, err := ParseControlNumber(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseControlNumber_WithSubsequence(t *testing.T) {
	input := "FRAGO-20-001-01"

	n := 1
	expected := ControlNumber{
		Class:        FRAGO,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  &n,
	}

	result, err := ParseControlNumber(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseControlNumber_ParsePrettyPrinted(t *testing.T) {
	input := "FRAGO--20--001--01"

	_, err := ParseControlNumber(input)

	assert.Error(t, err)
}

func TestControlNumber_String_WithoutSubsequence(t *testing.T) {
	input := ControlNumber{
		Class:        MR,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  nil,
	}

	expected := "MR-20-001"

	result := input.String()

	assert.Equal(t, expected, result)
}

func TestControlNumber_String_WithSubsequence(t *testing.T) {
	n := 1
	input := ControlNumber{
		Class:        FRAGO,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  &n,
	}

	expected := "FRAGO-20-001-01"

	result := input.String()

	assert.Equal(t, expected, result)
}

func TestControlNumber_PrettyPrint_WithoutSubsequence(t *testing.T) {
	input := ControlNumber{
		Class:        MR,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  nil,
	}

	expected := "MR--20--001"

	result := input.PrettyPrint()

	assert.Equal(t, expected, result)
}

func TestControlNumber_PrettyPrint_WithSubsequence(t *testing.T) {
	n := 1
	input := ControlNumber{
		Class:        FRAGO,
		Year:         2020,
		MainSequence: 1,
		SubSequence:  &n,
	}

	expected := "FRAGO--20--001--01"

	result := input.PrettyPrint()

	assert.Equal(t, expected, result)
}
