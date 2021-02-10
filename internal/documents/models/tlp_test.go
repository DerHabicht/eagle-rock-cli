// TODO: Refactor this to eagle-rock-api/pkg
package models

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestParseFullTlpString_SimpleTlp(t *testing.T) {
	test := "TLP:GREEN"

	resultTlp, resultCaveats, err := ParseFullTlpString(test)

	assert.NoError(t, err)
	assert.Equal(t, GREEN, resultTlp)
	assert.Nil(t, resultCaveats)
}

func TestParseFullTlpString_TlpWithCaveats(t *testing.T) {
	test := "TLP:RED//FAMILY/FRIENDS"

	resultTlp, resultCaveats, err := ParseFullTlpString(test)

	assert.NoError(t, err)
	assert.Equal(t, RED, resultTlp)
	assert.Equal(t, []string{"FAMILY", "FRIENDS"}, resultCaveats)
}

func TestTlp_MarshalYAML(t *testing.T) {
	expected := "testTlp: RED\n"

	test := struct {
		TestTlp Tlp `yaml:"testTlp"`
	}{
		TestTlp: RED,
	}

	result, err := yaml.Marshal(test)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}

func TestTlp_UnmarshalYAML(t *testing.T) {
	expected := struct {
		TestTlp Tlp `yaml:"testTlp"`
	}{
		TestTlp: RED,
	}

	test := `testTlp: RED`
	result := struct {
		TestTlp Tlp `yaml:"testTlp"`
	}{}

	err := yaml.Unmarshal([]byte(test), &result)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
