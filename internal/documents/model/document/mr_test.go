package document

import (
	"github.com/derhabicht/eagle-rock-lib/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMrHeader_IsIHeader(t *testing.T) {
	var _ IHeader = (*MrHeader)(nil)
	assert.True(t, true)
}

func TestMr_IsIDocument(t *testing.T) {
	var _ IDocument = (*Mr)(nil)
	assert.True(t, true)
}

func TestMrHeader_HeaderFieldMap_NonStandardsMrDoesNotHaveStatusFields(t *testing.T) {
	inputTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	inputDate, err := lib.ParseDate("2021-09-27")

	input := MrHeader{
		MemoHeader: MemoHeader{
			Logo:    HeaderLogo{
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
		Track:   ADVISORY,
		Subject: "subject",
		StatusHistory: MrStatusHistory{
			Registered: MrStatus{
				Date: inputDate,
			},
			Published: &MrStatus{
				Date: inputDate,
			},
		},
	}

	result := input.HeaderFieldMap()

	assert.Contains(t, result, "TRACK")
	assert.Contains(t, result, "SUBJECT")
}

func TestMrHeader_HeaderFieldMap_StandardsMrHasAllFields(t *testing.T) {
	inputTlp, err := lib.ParseTlp("TLP:RED//FOO/BAR")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	inputDate, err := lib.ParseDate("2021-09-27")

	input := MrHeader{
		MemoHeader: MemoHeader{
			Logo:    HeaderLogo{
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
		Track:   STANDARDS,
		Subject: "subject",
		StatusHistory: MrStatusHistory{
			Registered: MrStatus{
				Date: inputDate,
			},
			Published: &MrStatus{
				Date: inputDate,
			},
		},
	}

	result := input.HeaderFieldMap()

	assert.Contains(t, result, "TRACK")
	assert.Contains(t, result, "SUBJECT")
	assert.Contains(t, result, "ADOPTED")
	assert.Contains(t, result, "REJECTED")
}
