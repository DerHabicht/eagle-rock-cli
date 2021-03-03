package document

import (
	"github.com/pkg/errors"
	"strings"
)

type MrTrack int

const (
	STANDARDS MrTrack = iota + 1
	PROGRAM
	PROJECT
	ADVISORY
)

func ParseMrTrack(s string) (MrTrack, error) {
	switch strings.ToLower(s) {
	case "standards":
		return STANDARDS, nil
	case "program":
		return PROGRAM, nil
	case "project":
		return PROJECT, nil
	case "advisory":
		return ADVISORY, nil
	default:
		return -1, errors.Errorf("%s is not a valid MR track", s)
	}
}

func (mt MrTrack) String() string {
	switch mt {
	case STANDARDS:
		return "STANDARDS"
	case PROGRAM:
		return "PROGRAM"
	case PROJECT:
		return "PROJECT"
	case ADVISORY:
		return "ADVISORY"
	default:
		panic(errors.Errorf("%d is not a valid MR track", mt))
	}
}

func (mt MrTrack) MarshalYAML() (interface{}, error) {
	return mt.String(), nil
}

func (mt *MrTrack) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	temp, err := ParseMrTrack(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*mt = temp
	return nil
}
