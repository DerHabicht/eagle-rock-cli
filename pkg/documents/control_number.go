package documents

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type ControlNumber struct {
	Class        ControlNumberClass
	Year         int
	MainSequence int
	SubSequence  *int
}

func ParseControlNumber(s string) (ControlNumber, error) {
	cn := strings.Split(s, "-")

	if len(cn) < 3 {
		return ControlNumber{}, errors.Errorf("%s is not a valid control number", s)
	}

	class, err := ParseControlNumberClass(cn[0])
	if err != nil {
		return ControlNumber{}, errors.WithStack(err)
	}

	year, err := strconv.Atoi(cn[1])
	if err != nil {
		return ControlNumber{}, errors.WithMessage(err, "failed to parse control number year")
	}

	mainSeq, err := strconv.Atoi(cn[2])
	if err != nil {
		return ControlNumber{}, errors.WithMessage(err, "failed to parse the control number's main sequence")
	}

	var subSeq *int
	if len(cn) == 4 {
		n, err := strconv.Atoi(cn[3])
		if err != nil {
			return ControlNumber{}, errors.WithMessage(err, "failed to parse the control number's subsequence")
		}
		subSeq = &n
	}

	return ControlNumber{
			Class:        class,
			Year:         year + 2000,
			MainSequence: mainSeq,
			SubSequence:  subSeq,
		},
		nil
}

func (cn ControlNumber) String() string {
	if cn.SubSequence == nil {
		return fmt.Sprintf(
			"%s-%02d-%03d",
			cn.Class,
			cn.Year-2000,
			cn.MainSequence,
		)
	} else {
		return fmt.Sprintf(
			"%s-%02d-%03d-%02d",
			cn.Class,
			cn.Year-2000,
			cn.MainSequence,
			*cn.SubSequence,
		)
	}
}

func (cn ControlNumber) PrettyPrint() string {
	if cn.SubSequence == nil {
		return fmt.Sprintf(
			"%s--%02d--%03d",
			cn.Class,
			cn.Year-2000,
			cn.MainSequence,
		)
	} else {
		return fmt.Sprintf(
			"%s--%02d--%03d--%02d",
			cn.Class,
			cn.Year-2000,
			cn.MainSequence,
			*cn.SubSequence,
		)
	}
}

func (cn ControlNumber) MarshalYAML() (interface{}, error) {
	return cn.String(), nil
}

func (cn *ControlNumber) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := ParseControlNumber(buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*cn = t
	return nil
}
