// TODO: Refactor this to eagle-rock-api/pkg
package models

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type Tlp int

const (
	RED Tlp = iota + 1
	AMBER
	GREEN
	WHITE
)

func (t Tlp) String() string {
	switch t {
	case RED:
		return "RED"
	case AMBER:
		return "AMBER"
	case GREEN:
		return "GREEN"
	case WHITE:
		return "WHITE"
	default:
		panic(fmt.Sprintf("%d is not a valid Tlp value", t))
	}
}

func ParseTlp(t string) (Tlp, error) {
	switch strings.ToLower(t) {
	case "red":
		return RED, nil
	case "amber":
		return AMBER, nil
	case "green":
		return GREEN, nil
	case "white":
		return WHITE, nil
	default:
		return 0, errors.Errorf("%s cannot be parsed to a valid Tlp value", t)
	}
}
func BuildFullTlpString(tlp Tlp, caveats []string) string {
	if len(caveats) > 0 {
		return fmt.Sprintf(
			"TLP:%s//%s",
			tlp.String(),
			strings.Join(caveats, "/"),
		)
	}

	return fmt.Sprintf("TLP:%s", tlp.String())
}

func ParseFullTlpString(s string) (Tlp, []string, error) {
	var caveats []string

	pieces := strings.Split(s[4:], "//")

	if len(pieces) > 1 {
		caveats = strings.Split(pieces[1], "/")
	}

	tlp, err := ParseTlp(pieces[0])
	if err != nil {
		return 0, nil, errors.WithStack(err)
	}

	return tlp, caveats, nil
}

func (t Tlp) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}

func (t *Tlp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	tlp, err := ParseTlp(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*t = tlp
	return nil
}
