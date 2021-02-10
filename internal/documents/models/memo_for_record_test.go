package models

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestMemoForRecordTrack_MarshalYAML(t *testing.T) {
	expected := "testTrack: STANDARDS\n"

	test := struct {
		TestTrack MemoForRecordTrack `yaml:"testTrack"`
	}{
		TestTrack: STANDARDS,
	}

	result, err := yaml.Marshal(test)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}

func TestMemoForRecordTrack_UnmarshalYAML(t *testing.T) {
	expected := struct {
		TestTrack MemoForRecordTrack `yaml:"testTrack"`
	}{
		TestTrack: STANDARDS,
	}

	test := `testTrack: STANDARDS`
	result := struct {
		TestTrack MemoForRecordTrack `yaml:"testTrack"`
	}{}

	err := yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
