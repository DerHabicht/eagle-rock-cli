// TODO: Write tests
package documents

import (
	"github.com/pkg/errors"
	"strings"
)

type ControlNumberClass int

const (
	MR ControlNumberClass = iota + 1
	WARNO
	OPORD
	FRAGO
)

func ParseControlNumberClass(s string) (ControlNumberClass, error) {
	switch strings.ToLower(s) {
	case "mr":
		return MR, nil
	case "warno":
		return WARNO, nil
	case "opord":
		return OPORD, nil
	case "frago":
		return FRAGO, nil
	default:
		return -1, errors.Errorf("%s is not a valid control number class", s)
	}
}

func (cnc ControlNumberClass) String() string {
	switch cnc {
	case MR:
		return "MR"
	case WARNO:
		return "WARNO"
	case OPORD:
		return "OPORD"
	case FRAGO:
		return "FRAGO"
	default:
		panic(errors.Errorf("%d is not a valid control number class", cnc))
	}
}
