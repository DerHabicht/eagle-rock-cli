// TODO: Refactor this to eagle-rock-api/pkg
package models

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
