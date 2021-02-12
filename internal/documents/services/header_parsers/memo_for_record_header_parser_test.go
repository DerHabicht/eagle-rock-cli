package header_parsers

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoForRecordHeaderParser_ParseHeader_MemoForRecord(t *testing.T) {
	test := `---
tlp:            RED
caveats:
  - FAMILY
  - FRIENDS
control_number: MR-00-000
date:           2020-09-27
memo_for:       RECORD---STANDARDS TRACK
memo_from:      Robert Hawk
subject:        Test Memorandum
attachments:
  - Thing 1
  - Thing 2
cc:
  - Robert Hawk
...

####
This is a test memorandum.

####
For use in testing things.

####
By Testificates.
`
	expectedDate, err := time.Parse("2006-01-02", "2020-09-27")
	if err != nil {
		panic(err)
	}
	expectedHeader := models.MemoForRecordHeader{
		Tlp:           models.RED,
		Caveats:       []string{"FAMILY", "FRIENDS"},
		ControlNumber: "MR-00-000",
		Date:          &models.Date{Time: expectedDate},
		MemoFor:       "RECORD---STANDARDS TRACK",
		MemoFrom:      "Robert Hawk",
		Subject:       "Test Memorandum",
		Attachments:   []string{"Thing 1", "Thing 2"},
		Cc:            []string{"Robert Hawk"},
	}

	expectedBody := `

####
This is a test memorandum.

####
For use in testing things.

####
By Testificates.
`

	headerParser := NewMemoForRecordHeaderParser()

	resultHeader, resultBody, err := headerParser.ParseHeader([]byte(test))

	assert.NoError(t, err)
	assert.IsType(t, models.MemoForRecordHeader{}, resultHeader)
	assert.Equal(t, expectedHeader, resultHeader)
	assert.Equal(t, []byte(expectedBody), resultBody)

}
