package document

import (
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarnoHeader_IsIHeader(t *testing.T) {
	var _ IHeader = (*WarnoHeader)(nil)
	assert.True(t, true)
}

func TestWarnoSignature_IsISignature(t *testing.T) {
	var _ ISignature = (*WarnoSignature)(nil)
	assert.True(t, true)
}

func TestWarno_IsIDocument(t *testing.T) {
	var _ IDocument = (*Warno)(nil)
	assert.True(t, true)
}

func TestParseWarnoHeader_ValidHeader(t *testing.T) {
	input := `TLP:RED//FOO/BAR
20210209T2130Z THUS//HQ
WARNO-21-001 ACES-20-006-S01
TZ UTC-0700 IC ROBERT HAWK`

	expectedTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedDtg, err := lib.ParseDtg("20210209T2130Z")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expected := WarnoHeader{
		Tlp:      expectedTlp,
		DateTime: expectedDtg,
		Issuer:   "THUS//HQ",
		ControlNumber: lib.ControlNumber{
			Class:        lib.WARNO,
			Year:         2021,
			MainSequence: 1,
		},
		MissionNumber:     "ACES-20-006-S01",
		TimeZone:          "UTC-0700",
		IncidentCommander: "ROBERT HAWK",
	}

	result, err := ParseWarnoHeader(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestWarnoHeader_String(t *testing.T) {
	inputTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	inputDtg, err := lib.ParseDtg("20210209T2130Z")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	input := WarnoHeader{
		Tlp:      inputTlp,
		DateTime: inputDtg,
		Issuer:   "THUS//HQ",
		ControlNumber: lib.ControlNumber{
			Class:        lib.WARNO,
			Year:         2021,
			MainSequence: 1,
		},
		MissionNumber:     "ACES-20-006-S01",
		TimeZone:          "UTC-0700",
		IncidentCommander: "ROBERT HAWK",
	}

	expected := `TLP:RED//FOO/BAR
20210209T2130Z THUS//HQ
WARNO-21-001 ACES-20-006-S01
TZ UTC-0700 IC ROBERT HAWK`

	result := input.String()

	assert.Equal(t, expected, result)
}
