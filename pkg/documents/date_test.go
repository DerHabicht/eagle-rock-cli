// TODO: Refactor this to eagle-rock-api/pkg
package documents

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
	"time"
)

func TestDate_MarshalYAML(t *testing.T) {
	// TODO: Figure out how to marshal a date without the enclosing quotes
	testTime, err := time.Parse("2006-01-02", "1988-09-27")
	if err != nil {
		panic(err)
	}
	test := struct {
		TestDate Date `yaml:"date"`
	}{
		TestDate: Date{testTime},
	}

	result, err := yaml.Marshal(&test)

	assert.NoError(t, err)
	assert.Equal(t, "date: \"1988-09-27\"\n", string(result))
}

func TestDate_UnmarshalYAML_NotNull(t *testing.T) {
	expected, err := time.Parse("2006-01-02", "1988-09-27")
	if err != nil {
		panic(err)
	}

	test := "date: \"1988-09-27\""

	result := struct {
		TestDtg Date `yaml:"date"`
	}{}

	err = yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Equal(t, Date{expected}, result.TestDtg)
}

func TestDate_UnmarshalYAML_Null(t *testing.T) {
	test := "date: null"

	result := struct {
		TestDtg *Date `yaml:"date"`
	}{}

	err := yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Nil(t, result.TestDtg)
}

func TestParseDtg_WillParseLong(t *testing.T) {
	input := "2012-09-27T16:14Z"

	dtg, err := time.Parse("2006-01-02T15:04Z", input)
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expected := Dtg{dtg}

	result, err := ParseDtg(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseDtg_WillParseShort(t *testing.T) {
	input := "20120927T1614Z"

	dtg, err := time.Parse("20060102T1504Z", input)
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expected := Dtg{dtg}

	result, err := ParseDtg(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseDtg_WillParseExplicitTimeZoneLong(t *testing.T) {
	input := "2012-09-27T16:14-07:00"

	dtg, err := time.Parse("2006-01-02T15:04-07:00", input)
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expected := Dtg{dtg}

	result, err := ParseDtg(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestParseDtg_WillParseExplicitTimeZoneShort(t *testing.T) {
	input := "20120927T1614-0700"

	dtg, err := time.Parse("20060102T1504-0700", input)
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expected := Dtg{dtg}

	result, err := ParseDtg(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDtg_FormatLong_WillFormatZForUTC(t *testing.T) {
	inputTime, err := time.Parse("2006-01-02T15:04-07:00", "2012-09-27T16:14+00:00")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	input := Dtg{inputTime}

	expected := "2012-09-27T16:14Z"

	result := input.FormatLong()

	assert.Equal(t, expected, result)
}

func TestDtg_FormatLong_WillFormatExplicitTimeZone(t *testing.T) {
	inputTime, err := time.Parse("2006-01-02T15:04-07:00", "2012-09-27T16:14-07:00")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	input := Dtg{inputTime}

	expected := "2012-09-27T16:14-07:00"

	result := input.FormatLong()

	assert.Equal(t, expected, result)
}

func TestDtg_FormatShort_WillFormatZForUTC(t *testing.T) {
	inputTime, err := time.Parse("2006-01-02T15:04-07:00", "2012-09-27T16:14+00:00")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	input := Dtg{inputTime}

	expected := "20120927T1614Z"

	result := input.FormatShort()

	assert.Equal(t, expected, result)
}

func TestDtg_FormatShort_WillFormatExplicitTimeZone(t *testing.T) {
	inputTime, err := time.Parse("2006-01-02T15:04-07:00", "2012-09-27T16:14-07:00")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	input := Dtg{inputTime}

	expected := "20120927T1614-0700"

	result := input.FormatShort()

	assert.Equal(t, expected, result)
}

func TestDtg_MarshalYAML(t *testing.T) {
	// TODO: Figure out how to marshal a date without the enclosing quotes
	testTime, err := time.Parse("2006-01-02T15:04", "1988-09-27T16:42")
	if err != nil {
		panic(err)
	}
	test := struct {
		TestDtg Dtg `yaml:"dtg"`
	}{
		TestDtg: Dtg{testTime},
	}

	result, err := yaml.Marshal(&test)

	assert.NoError(t, err)
	assert.Equal(t, "dtg: 1988-09-27T16:42Z\n", string(result))
}

func TestDtg_UnmarshalYAML_NotNull(t *testing.T) {
	expected, err := time.Parse("2006-01-02T15:04Z", "1988-09-27T16:42Z")
	if err != nil {
		panic(err)
	}

	test := "dtg: 1988-09-27T16:42Z"

	result := struct {
		TestDtg Dtg `yaml:"dtg"`
	}{}

	err = yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Equal(t, Dtg{expected}, result.TestDtg)
}

func TestDtg_UnmarshalYAML_Null(t *testing.T) {
	test := "dtg: null"

	result := struct {
		TestDtg *Dtg `yaml:"dtg"`
	}{}

	err := yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Nil(t, result.TestDtg)
}
