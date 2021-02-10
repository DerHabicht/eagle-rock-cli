package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWarnoHeader_Marshal(t *testing.T) {
	expected := `TLP:GREEN
20210209T2130Z THUS//HQ
WARNO-21-001 ACES-20-006-S01
TZ UTC-07 IC ROBERT HAWK`

	testTime, err := time.Parse("20060102T1504Z", "20210209T2130Z")
	if err != nil {
		panic(err)
	}

	test := WarnoHeader{
		Tlp:               GREEN,
		DateTime:          testTime,
		Issuer:            "THUS//HQ",
		ControlNumber:     "WARNO-21-001",
		Mission:           "ACES-20-006-S01",
		TimeZone:          "UTC-07",
		IncidentCommander: "Robert Hawk",
	}

	result, err := test.Marshal()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
func TestWarnoHeader_Unmarshal_SimpleTlp(t *testing.T) {
	expectedTime, err := time.Parse("20060102T1504Z", "20210209T2130Z")
	if err != nil {
		panic(err)
	}

	expected := WarnoHeader{
		Tlp:               GREEN,
		DateTime:          expectedTime,
		Issuer:            "THUS//HQ",
		ControlNumber:     "WARNO-21-001",
		Mission:           "ACES-20-006-S01",
		TimeZone:          "UTC-07",
		IncidentCommander: "Robert Hawk",
	}

	test := `
TLP:GREEN
20210209T2130Z THUS//HQ
WARNO-21-001 ACES-20-006-S01
TZ UTC-07 IC ROBERT HAWK
`

	result, err := UnmarshalWarnoHeader(test)

	assert.NoError(t, err)

	assert.Equal(t, expected.Tlp, result.Tlp)
	assert.Nil(t, expected.Caveats)
	assert.Equal(t, expected.DateTime, result.DateTime)
	assert.Equal(t, expected.Issuer, result.Issuer)
	assert.Equal(t, expected.ControlNumber, result.ControlNumber)
	assert.Equal(t, expected.Mission, result.Mission)
	assert.Equal(t, expected.TimeZone, result.TimeZone)
	assert.Equal(t, expected.IncidentCommander, result.IncidentCommander)
}

func TestWarnoHeader_Unmarshal_TlpWithCaveats(t *testing.T) {
	expectedTime, err := time.Parse("20060102T1504Z", "20210209T2130Z")
	if err != nil {
		panic(err)
	}

	expected := WarnoHeader{
		Tlp:               RED,
		Caveats:           []string{"FAMILY", "FRIENDS"},
		DateTime:          expectedTime,
		Issuer:            "THUS//HQ",
		ControlNumber:     "WARNO-21-001",
		Mission:           "ACES-20-006-S01",
		TimeZone:          "UTC-07",
		IncidentCommander: "Robert Hawk",
	}

	test := `
TLP:RED//FAMILY/FRIENDS
20210209T2130Z THUS//HQ
WARNO-21-001 ACES-20-006-S01
TZ UTC-07 IC ROBERT HAWK
`

	result, err := UnmarshalWarnoHeader(test)

	assert.NoError(t, err)

	assert.Equal(t, expected.Tlp, result.Tlp)
	assert.Equal(t, expected.DateTime, result.DateTime)
	assert.Equal(t, expected.Issuer, result.Issuer)
	assert.Equal(t, expected.ControlNumber, result.ControlNumber)
	assert.Equal(t, expected.Mission, result.Mission)
	assert.Equal(t, expected.TimeZone, result.TimeZone)
	assert.Equal(t, expected.IncidentCommander, result.IncidentCommander)
}
