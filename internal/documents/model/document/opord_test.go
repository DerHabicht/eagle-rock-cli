package document

import (
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpordHeader_IsIHeader(t *testing.T) {
	var _ IHeader = (*OpordHeader)(nil)
	assert.True(t, true)
}

func TestOpord_IsIDocument(t *testing.T) {
	var _ IDocument = (*Opord)(nil)
	assert.True(t, true)
}

func TestOpordHeader_HeaderFieldMap_HasAllFields(t *testing.T) {
	inputTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	inputDate, err := lib.ParseDate("2021-09-27")

	input := OpordHeader{
		MemoHeader: MemoHeader{
			Logo: HeaderLogo{
				Image: "logo",
				Scale: 1.0,
			},
			Address: "address",
			Tlp:     inputTlp,
			ControlNumber: lib.ControlNumber{
				Class:        lib.MR,
				Year:         2021,
				MainSequence: 1,
			},
			Date:        &inputDate,
			Attachments: nil,
			Cc:          nil,
		},
		MissionNumber:     "mission_number",
		TimeZone:          "UTC+0000",
		IncidentCommander: "IncidentCommander",
	}

	result := input.HeaderFieldMap()

	assert.Contains(t, result, "MISSION_NUMBER")
	assert.Contains(t, result, "TIMEZONE")
	assert.Contains(t, result, "INCIDENT_COMMANDER")
}

func TestOpordHeader_Title(t *testing.T) {
	inputTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	inputDate, err := lib.ParseDate("2021-09-27")

	input := OpordHeader{
		MemoHeader: MemoHeader{
			Logo: HeaderLogo{
				Image: "logo",
				Scale: 1.0,
			},
			Address: "address",
			Tlp:     inputTlp,
			ControlNumber: lib.ControlNumber{
				Class:        lib.MR,
				Year:         2021,
				MainSequence: 1,
			},
			Date:        &inputDate,
			Attachments: nil,
			Cc:          nil,
		},
		MissionNumber:     "mission_number",
		TimeZone:          "UTC+0000",
		IncidentCommander: "IncidentCommander",
	}

	result := input.DocumentTitle()

	assert.Equal(t, "mission_number", result)
}